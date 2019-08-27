/* ######################################################################
# File Name: models/message_task/search.go
# Author: Rain
# Main: jiayd163@163.com
# Created Time: 2019-02-12 14:32:43
####################################################################### */
package message_task

import (
	"github.com/ant-libs-go/util"
	"github.com/go-xorm/builder"
)

type search struct {
	MessageTask
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

func (this *search) Search() (r []*MessageTask, r2 map[int32]*MessageTask, err error) {
	query := New().Active().OrderBy("id DESC")

	if this.LastId > 0 {
		query.And(builder.Lt{"id": this.LastId}).Limit(int(this.Limit))
	} else if this.Limit == 0 {
		query.Limit(int(this.Limit), int(this.Offset))
	}
	if this.Id > 0 {
		query.And(builder.Eq{"id": this.Id})
	}
	if len(this.Ids) > 0 {
		query.And(builder.Eq{"id": this.Ids})
	}
	return query.Find()
}
