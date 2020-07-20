/* ######################################################################
# Author: (zfly1207@126.com)
# Created Time: 2018-08-29 08:17:31
# File Name: models.go
# Description:
####################################################################### */

package models

import (
	"encoding/gob"
	"fmt"
	"time"

	"message_sender/libs/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	types "gitlab.com/feichi/fcad_thrift/libs/go/fcmp_passport_types"
)

var (
	Orm      *xorm.EngineGroup
	pageSize = 10
)

func Init() {
	main, err := xorm.NewEngine("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&autocommit=true",
		config.Get().Db["main"].User,
		config.Get().Db["main"].Pawd,
		config.Get().Db["main"].Host,
		config.Get().Db["main"].Port,
		config.Get().Db["main"].Name))
	if err != nil {
		panic(fmt.Sprintf("main db connect errr: %s", err))
	}
	subordinate, err := xorm.NewEngine("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&autocommit=true",
		config.Get().Db["subordinate"].User,
		config.Get().Db["subordinate"].Pawd,
		config.Get().Db["subordinate"].Host,
		config.Get().Db["subordinate"].Port,
		config.Get().Db["subordinate"].Name))
	if err != nil {
		panic(fmt.Sprintf("subordinate db connect errr: %s", err))
	}
	Orm, err = xorm.NewEngineGroup(main, []*xorm.Engine{subordinate})
	Orm.SetConnMaxLifetime(3000)

	Orm.ShowSQL(config.Get().Basic.Debug)
	Orm.TZLocation, _ = time.LoadLocation("Asia/Shanghai")

	gob.Register(types.MpUser{})
}

type SendStatus int

const (
	SendStatusAwait    SendStatus = 1 // 待发送
	SendStatusProgress SendStatus = 2 // 发送中
	SendStatusSucc     SendStatus = 3 // 完成
	SendStatusFail     SendStatus = 4 // 失败
)
