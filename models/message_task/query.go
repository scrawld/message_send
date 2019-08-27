/* ######################################################################
# File Name: models/message_task/query.go
# Author: Rain
# Main: jiayd163@163.com
# Created Time: 2019-02-12 14:32:25
####################################################################### */
package message_task

import (
	"errors"
	"fcmp_message_sender/models"
	"time"

	"github.com/go-xorm/builder"
	"gitlab.com/feichi/fcad_thrift/libs/go/enums"
)

func (this *MessageTaskQuery) Active() *MessageTaskQuery {
	return this.And(builder.Eq{"status": enums.InfoStatus_Normal})
}

/* query */
func (this *MessageTaskQuery) GetById(id int32) (r *MessageTask, err error) {
	if id == 0 {
		return
	}
	return this.Active().Where(builder.Eq{"id": id}).Get()
}

func (this *MessageTaskQuery) GetByIds(ids []int32) (r []*MessageTask, r2 map[int32]*MessageTask, err error) {
	if len(ids) == 0 {
		return
	}
	return this.Active().Where(builder.Eq{"id": ids}).Find()
}

// 获取待发送任务
func (this *MessageTaskQuery) GetAwaitTask() (r *MessageTask, err error) {
	r, err = this.Active().
		Where(builder.Lte{"start_at": time.Now().Unix()}.And(builder.Eq{"send_status": models.SendStatusAwait})).
		Get()
	if err != nil || r == nil {
		return
	}
	this.Id = r.Id
	this.SendStatus = models.SendStatusProgress
	err = this.Update()
	return
}

// 获取用户ids
func (this *MessageTaskQuery) GetUserIdBySql(query string) (r []int32, err error) {
	if len(query) == 0 {
		return
	}
	err = this.Session().Sql(query).Find(&r)
	return
}

/* create */
func (this *MessageTaskQuery) Create() (err error) {
	_, err = this.Session().Insert(&this.MessageTask)
	return
}

/* update */
func (this *MessageTaskQuery) Update() (err error) {
	if this.Id == 0 {
		return errors.New("id not set")
	}
	_, err = this.Where(builder.Eq{"id": this.Id}).Session().Update(&this.MessageTask)
	return
}
