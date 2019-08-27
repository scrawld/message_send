/* ######################################################################
# File Name: models/message_task/message_task.go
# Author: Rain
# Main: jiayd163@163.com
# Created Time: 2019-02-12 14:08:28
####################################################################### */
package message_task

import (
	"fcmp_message_sender/models"
	"strings"

	"github.com/ant-libs-go/util"
	"github.com/go-xorm/builder"
	"github.com/go-xorm/xorm"
	"gitlab.com/feichi/fcad_thrift/libs/go/enums"
)

type MessageTask struct {
	Id            int32 `xorm:"pk autoincr"`
	MediaId       int32
	TemplateId    string
	TemplateTitle string
	Name          string
	UserIds       string
	Sql           string
	Message       string
	ActionPath    string
	StartAt       int
	SendStatus    models.SendStatus
	SuccCount     int
	FailCount     int
	Status        enums.InfoStatus
	CreatedAt     int64 `xorm:"created"`
	UpdatedAt     int64 `xorm:"updated"`
}

func (this *MessageTask) TableName() string {
	return "message_task"
}

type MessageTaskQuery struct {
	MessageTask `xorm:"-"`
	session     *xorm.Session `xorm:"-"`
	isNewRecord bool          `xorm:"-"`
}

func New() *MessageTaskQuery {
	o := &MessageTaskQuery{}
	o.isNewRecord = true
	return o
}

func (this *MessageTaskQuery) Orm() (r *xorm.EngineGroup) {
	return models.Orm
}

func (this *MessageTaskQuery) Session() (r *xorm.Session) {
	if this.session == nil {
		this.session = this.Orm().NewSession()
	}
	return this.session
}

func (this *MessageTaskQuery) Load(inp interface{}, excludes ...string) *MessageTaskQuery {
	util.Assign(inp, this, excludes...)
	return this
}

func (this *MessageTaskQuery) SQL(query string, params ...interface{}) *MessageTaskQuery {
	this.Session().Sql(query, params...)
	return this
}

/* cond */
func (this *MessageTaskQuery) Where(cond builder.Cond) *MessageTaskQuery {
	this.Session().Where(cond)
	return this
}

func (this *MessageTaskQuery) And(cond builder.Cond) *MessageTaskQuery {
	this.Session().And(cond)
	return this
}

func (this *MessageTaskQuery) Or(cond builder.Cond) *MessageTaskQuery {
	this.Session().Or(cond)
	return this
}

/* misc */
func (this *MessageTaskQuery) Cols(cols ...string) *MessageTaskQuery {
	this.Session().Cols(cols...)
	return this
}

func (this *MessageTaskQuery) Select(str string) *MessageTaskQuery {
	this.Session().Select(str)
	return this
}

func (this *MessageTaskQuery) OrderBy(orders ...string) *MessageTaskQuery {
	this.Session().OrderBy(strings.Join(orders, ", "))
	return this
}

func (this *MessageTaskQuery) GroupBy(groups ...string) *MessageTaskQuery {
	this.Session().GroupBy(strings.Join(groups, ", "))
	return this
}

func (this *MessageTaskQuery) Limit(limit int, offset ...int) *MessageTaskQuery {
	this.Session().Limit(limit, offset...)
	return this
}

/* query */
func (this *MessageTaskQuery) Get() (r *MessageTask, err error) {
	r = &MessageTask{}
	if has, err := this.Session().Get(r); has == false || err != nil {
		return nil, err
	}
	this.isNewRecord = false
	return
}

func (this *MessageTaskQuery) Find() (r []*MessageTask, r2 map[int32]*MessageTask, err error) {
	err = this.Session().Find(&r)
	r2 = map[int32]*MessageTask{}
	for _, m := range r {
		r2[m.Id] = m
	}
	return
}

func (this *MessageTaskQuery) Count() (r int64, err error) {
	return this.Session().Count(&MessageTaskQuery{})
}

func (this *MessageTaskQuery) Exist() (r bool, err error) {
	return this.Session().Exist(&MessageTaskQuery{})
}
