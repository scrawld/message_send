/* ######################################################################
# File Name: work_handler/task_handler.go
# Author: Rain
# Main: jiayd163@163.com
# Created Time: 2019-02-12 14:59:14
####################################################################### */
package work_handler

import (
	"encoding/json"
	"fcmp_message_sender/libs/passport_server"
	"fcmp_message_sender/libs/utils"
	"fcmp_message_sender/models"
	"fcmp_message_sender/models/message_records"
	"fcmp_message_sender/models/message_task"
	"strings"

	"github.com/ant-libs-go/util/logs"
	uuid "github.com/satori/go.uuid"
	types "gitlab.com/feichi/fcad_thrift/libs/go/fcmp_passport_types"
)

type TaskHandler struct {
	Name string
	log  *logs.SessLog
}

func NewTaskHandler() *TaskHandler {
	o := &TaskHandler{}
	o.Name = "TASK_HANDLER"
	o.log = logs.New(uuid.NewV4().String())
	return o
}

func (this *TaskHandler) Run() {
	this.log.Infof("Recall#%s...", this.Name)
	defer this.log.Infof("Recall#%s finish", this.Name)
	task, err := message_task.New().GetAwaitTask()
	if err != nil {
		this.log.Warnf("Get task error: %s", err)
		return
	}
	if task == nil {
		this.log.Infof("Task is empty")
		return
	}
	this.writeRecord(task)
}

func (this *TaskHandler) writeRecord(task *message_task.MessageTask) {
	taskQuery := message_task.New()
	defer func() {
		taskQuery.Id = task.Id
		if err := taskQuery.Update(); err != nil {
			this.log.Warnf("Update task SendStatus error: %s", err)
			return
		}
	}()

	userIds, err := getUserIdsByTask(task)
	if err != nil {
		taskQuery.SendStatus = models.SendStatusFail
		this.log.Warnf("Get userId error: %s", err)
		return
	}
	for _, uId := range userIds {
		u, err := passport_server.GetUserById(uId)
		if err != nil {
			this.log.Warnf("Get user error: %s - userId: %d", err, uId)
			continue
		}
		if u == nil {
			this.log.Warnf("Get user is empty: userId: %d", uId)
			continue
		}
		var recordsQuery = message_records.New().Load(task, "Id", "Message", "SendStatus")
		recordsQuery.TaskId = task.Id
		recordsQuery.UserId = u.Id
		recordsQuery.WxOpenid = u.WxOpenid
		recordsQuery.Message = getMsgByUser(u, task.Message)
		recordsQuery.SendStatus = models.SendStatusAwait
		err = recordsQuery.Create()
		if err != nil {
			this.log.Warnf("Create record error: %s - userId: %d", err, uId)
			continue
		}
	}
	taskQuery.SendStatus = models.SendStatusSucc
	return

}

func getUserIdsByTask(task *message_task.MessageTask) (r []int32, err error) {
	if len(task.Sql) > 0 {
		var userIdsFromSql []int32
		userIdsFromSql, err = message_task.New().GetUserIdBySql(task.Sql)
		if err != nil {
			return
		}
		r = append(r, userIdsFromSql...)
	}
	if len(task.UserIds) > 0 {
		var userIdsFromJson []int32
		err = json.Unmarshal([]byte(task.UserIds), &userIdsFromJson)
		if err != nil {
			return
		}
		r = append(r, userIdsFromJson...)
	}
	r = utils.SliceUniqueInt32(r)
	return
}

func getMsgByUser(user *types.MpUser, msg string) (r string) {
	var customNickNamePlace = "{{custom_nick_name}}"
	var nickNamePlace = "{{nick_name}}"
	r = msg
	r = strings.Replace(r, customNickNamePlace, user.CustomNickname, -1)
	r = strings.Replace(r, nickNamePlace, user.Nickname, -1)
	return
}
