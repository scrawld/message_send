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
	master, err := xorm.NewEngine("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&autocommit=true",
		config.Get().Db["master"].User,
		config.Get().Db["master"].Pawd,
		config.Get().Db["master"].Host,
		config.Get().Db["master"].Port,
		config.Get().Db["master"].Name))
	if err != nil {
		panic(fmt.Sprintf("master db connect errr: %s", err))
	}
	slave, err := xorm.NewEngine("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&autocommit=true",
		config.Get().Db["slave"].User,
		config.Get().Db["slave"].Pawd,
		config.Get().Db["slave"].Host,
		config.Get().Db["slave"].Port,
		config.Get().Db["slave"].Name))
	if err != nil {
		panic(fmt.Sprintf("slave db connect errr: %s", err))
	}
	Orm, err = xorm.NewEngineGroup(master, []*xorm.Engine{slave})
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
