/* ######################################################################
# File Name: models/message_records/query.go
# Author: Rain
# Main: jiayd163@163.com
# Created Time: 2019-02-12 14:33:31
####################################################################### */
package message_records

import (
	"errors"
	"fcmp_message_sender/models"

	"github.com/go-xorm/builder"
	"gitlab.com/feichi/fcad_thrift/libs/go/enums"
)

func (this *MessageRecordsQuery) Active() *MessageRecordsQuery {
	return this.And(builder.Eq{"status": enums.InfoStatus_Normal})
}

/* query */
func (this *MessageRecordsQuery) GetById(id int32) (r *MessageRecords, err error) {
	if id == 0 {
		return
	}
	return this.Active().Where(builder.Eq{"id": id}).Get()
}

func (this *MessageRecordsQuery) GetByIds(ids []int32) (r []*MessageRecords, r2 map[int32]*MessageRecords, err error) {
	if len(ids) == 0 {
		return
	}
	return this.Active().Where(builder.Eq{"id": ids}).Find()
}

func (this *MessageRecordsQuery) RetryRecords() (err error) {
	sql := "UPDATE message_records SET send_status=1, is_retry=0, retry_count=retry_count+1 WHERE is_retry=1 AND retry_count<3"
	_, err = this.Orm().Exec(sql)
	return err
}

/* create */
func (this *MessageRecordsQuery) Create() (err error) {
	_, err = this.Session().Insert(&this.MessageRecords)
	return
}

/* update */
func (this *MessageRecordsQuery) Update() (err error) {
	if this.Id == 0 {
		return errors.New("id not set")
	}
	_, err = this.Cols("send_status", "response", "is_retry", "retry_count").Where(builder.Eq{"id": this.Id}).Session().Update(&this.MessageRecords)
	return
}

func (this *MessageRecordsQuery) UpdateSetSendStatus(record *MessageRecords, sendStatus models.SendStatus) (err error) {
	record.SendStatus = sendStatus
	_, err = this.Load(record).Cols("send_status").Where(builder.Eq{"id": record.Id}).Session().Update(&this.MessageRecords)
	return
}
