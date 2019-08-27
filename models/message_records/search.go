/* ######################################################################
# File Name: models/message_records/search.go
# Author: Rain
# Main: jiayd163@163.com
# Created Time: 2019-02-12 14:33:18
####################################################################### */
package message_records

import (
	"message_sender/models"

	"github.com/ant-libs-go/util"
	"github.com/go-xorm/builder"
)

type search struct {
	MessageRecords
	Ids    []int32
	LastId int32
	Limit  int32
	Offset int32
}

func NewSearch() *search {
	o := &search{}
	return o
}

func (this *search) Load(inp interface{}, excludes ...string) *search {
	util.Assign(inp, this, excludes...)
	return this
}

func (this *search) Search() (r []*MessageRecords, r2 map[int32]*MessageRecords, err error) {
	query := New().Active().OrderBy("id ASC")

	if this.LastId > 0 {
		query.And(builder.Lt{"id": this.LastId}).Limit(int(this.Limit))
	} else if this.Limit > 0 {
		query.Limit(int(this.Limit), int(this.Offset))
	}
	if this.Id > 0 {
		query.And(builder.Eq{"id": this.Id})
	}
	if len(this.Ids) > 0 {
		query.And(builder.Eq{"id": this.Ids})
	}
	if this.SendStatus > 0 {
		query.And(builder.Eq{"send_status": this.SendStatus})
	}
	return query.Find()
}

func (this *search) SearchBySendStatusAndLimit(sendStatus models.SendStatus, limit int32) (r []*MessageRecords, r2 map[int32]*MessageRecords, err error) {
	this.SendStatus = sendStatus
	this.Limit = limit
	return this.Search()
}
