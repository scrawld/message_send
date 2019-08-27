/* ######################################################################
# File Name: models/message_records/message_records.go
# Author: Rain
# Main: jiayd163@163.com
# Created Time: 2019-02-12 14:33:04
####################################################################### */
package message_records

import (
	"encoding/json"
	"fcmp_message_sender/models"
	"strings"

	"github.com/ant-libs-go/util"
	"github.com/go-xorm/builder"
	"github.com/go-xorm/xorm"
	"gitlab.com/feichi/fcad_thrift/libs/go/enums"
)

type MessageRecords struct {
	Id         int32 `xorm:"pk autoincr"`
	TaskId     int32
	MediaId    int32
	UserId     int32
	WxOpenid   string
	TemplateId string
	Message    string
	ActionPath string
	SendStatus models.SendStatus
	Response   string
	IsRetry    int
	RetryCount int
	Status     enums.InfoStatus
	CreatedAt  int64 `xorm:"created"`
	UpdatedAt  int64 `xorm:"updated"`
}

func (this *MessageRecords) TableName() string {
	return "message_records"
}

func (this *MessageRecords) DecodeMessage() (r map[string]interface{}, err error) {
	r = map[string]interface{}{}
	err = json.Unmarshal([]byte(this.Message), &r)
	return
}

type MessageRecordsQuery struct {
	MessageRecords `xorm:"-"`
	session        *xorm.Session `xorm:"-"`
	isNewRecord    bool          `xorm:"-"`
}

func New() *MessageRecordsQuery {
	o := &MessageRecordsQuery{}
	o.isNewRecord = true
	return o
}

func (this *MessageRecordsQuery) Orm() (r *xorm.EngineGroup) {
	return models.Orm
}

func (this *MessageRecordsQuery) Session() (r *xorm.Session) {
	if this.session == nil {
		this.session = this.Orm().NewSession()
	}
	return this.session
}

func (this *MessageRecordsQuery) Load(inp interface{}, excludes ...string) *MessageRecordsQuery {
	util.Assign(inp, this, excludes...)
	return this
}

func (this *MessageRecordsQuery) SQL(query string, params ...interface{}) *MessageRecordsQuery {
	this.Session().Sql(query, params...)
	return this
}

/* cond */
func (this *MessageRecordsQuery) Where(cond builder.Cond) *MessageRecordsQuery {
	this.Session().Where(cond)
	return this
}

func (this *MessageRecordsQuery) And(cond builder.Cond) *MessageRecordsQuery {
	this.Session().And(cond)
	return this
}

func (this *MessageRecordsQuery) Or(cond builder.Cond) *MessageRecordsQuery {
	this.Session().Or(cond)
	return this
}

/* misc */
func (this *MessageRecordsQuery) Cols(cols ...string) *MessageRecordsQuery {
	this.Session().Cols(cols...)
	return this
}

func (this *MessageRecordsQuery) Select(str string) *MessageRecordsQuery {
	this.Session().Select(str)
	return this
}

func (this *MessageRecordsQuery) OrderBy(orders ...string) *MessageRecordsQuery {
	this.Session().OrderBy(strings.Join(orders, ", "))
	return this
}

func (this *MessageRecordsQuery) GroupBy(groups ...string) *MessageRecordsQuery {
	this.Session().GroupBy(strings.Join(groups, ", "))
	return this
}

func (this *MessageRecordsQuery) Limit(limit int, offset ...int) *MessageRecordsQuery {
	this.Session().Limit(limit, offset...)
	return this
}

/* query */
func (this *MessageRecordsQuery) Get() (r *MessageRecords, err error) {
	r = &MessageRecords{}
	if has, err := this.Session().Get(r); has == false || err != nil {
		return nil, err
	}
	this.isNewRecord = false
	return
}

func (this *MessageRecordsQuery) Find() (r []*MessageRecords, r2 map[int32]*MessageRecords, err error) {
	err = this.Session().Find(&r)
	r2 = map[int32]*MessageRecords{}
	for _, m := range r {
		r2[m.Id] = m
	}
	return
}

func (this *MessageRecordsQuery) Count() (r int64, err error) {
	return this.Session().Count(&MessageRecordsQuery{})
}

func (this *MessageRecordsQuery) Exist() (r bool, err error) {
	return this.Session().Exist(&MessageRecordsQuery{})
}
