/* ######################################################################
# File Name: work_handler/sender_handler.go
# Author: Rain
# Main: jiayd163@163.com
# Created Time: 2019-02-13 14:44:53
####################################################################### */
package work_handler

import (
	"sync"
	"time"

	"message_sender/libs/config"
	"message_sender/libs/easywechat"
	"message_sender/libs/passport_server"
	"message_sender/libs/sendingerror"
	"message_sender/models"
	"message_sender/models/media"
	"message_sender/models/message_records"

	"github.com/ant-libs-go/util/logs"
	uuid "github.com/satori/go.uuid"
)

var (
	wg         = &sync.WaitGroup{}
	recordChan = make(chan *message_records.MessageRecords)
)

type MessageHandler struct {
	running bool
}

func NewMessageHandler() *MessageHandler {
	o := &MessageHandler{
		running: true,
	}
	return o
}

func (this *MessageHandler) Start() {
	defer wg.Wait()

	for i := 0; i < config.Get().Sender.RoutineNum; i++ {
		wg.Add(1)
		go func() { this.worker(); wg.Done() }()
	}

	go this.getRecord()

	go this.RetryRecords()
}

func (this *MessageHandler) Stop() {
	close(recordChan)
	this.running = false
}

func (this *MessageHandler) getRecord() {
	for this.running {
		log := logs.New(uuid.NewV4().String())
		records, _, err := message_records.NewSearch().SearchBySendStatusAndLimit(models.SendStatusAwait, int32(config.Get().Sender.Limit))
		if err != nil {
			log.Warnf("Get records error: %s", err)
			time.Sleep(time.Second * 5)
			continue
		}
		if records == nil {
			log.Infof("Records is empty")
			time.Sleep(time.Second * 5)
			continue
		}
		for _, r := range records {
			recordChan <- r
		}
		time.Sleep(time.Second)
	}
}

func (this *MessageHandler) worker() {
	for this.running {
		log := logs.New(uuid.NewV4().String())
		record := <-recordChan
		if record == nil {
			log.Infof("recordChan is close")
			continue
		}
		log.Infof("Start to work: %#v", record)
		if err := message_records.New().UpdateSetSendStatus(record, models.SendStatusProgress); err != nil {
			log.Warnf("Update records error(SendStatusProgress): %s", err)
			continue
		}
		if err := send(record); err != nil {
			log.Warnf("Send records error: %s", err)
			continue
		}
	}
}

func send(record *message_records.MessageRecords) (err error) {
	var sendingErr *sendingerror.SendingError
	var recordsQuery = message_records.New().Load(record)
	defer func() {
		if sendingErr != nil {
			recordsQuery.Load(sendingErr)
		}
		err = recordsQuery.Update()
	}()
	accessToken := media.GetAccessTokenById(record.MediaId)
	if len(record.WxOpenid) == 0 || len(accessToken) == 0 {
		sendingErr = sendingerror.BuildErr(0, "openid or access_token is empty")
		return
	}
	message, err := record.DecodeMessage()
	if err != nil {
		sendingErr = sendingerror.BuildErr(0, "json Unmarshal message error: %s", err)
		return
	}
	formId, err := passport_server.GetFormidByUserId(record.UserId)
	if err != nil {
		sendingErr = sendingerror.BuildErr(1, "Get formId error: %s - userId: %d", err, record.UserId)
		return
	}
	if len(formId) == 0 {
		sendingErr = sendingerror.BuildErr(0, "formId is empty")
		return
	}
	log := logs.New(uuid.NewV4().String())
	log.Infof("SendTemplateMessage: user_id: %d --- records_id: %d --- formId: %s", record.UserId, record.Id, formId)
	sendRes, err := easywechat.SendTemplateMessage(accessToken, record.WxOpenid, record.TemplateId, formId, record.ActionPath, message)
	if err != nil {
		sendingErr = sendingerror.BuildErr(sendRes.IsRetry, "Request template message error: %s", err)
		return
	}
	recordsQuery.SendStatus = models.SendStatusSucc
	return
}

func (this *MessageHandler) RetryRecords() {
	for this.running {
		log := logs.New(uuid.NewV4().String())
		err := message_records.New().RetryRecords()
		if err != nil {
			log.Warnf("Update records error: %s", err)
			return
		}
		time.Sleep(time.Second * 2)
	}
}
