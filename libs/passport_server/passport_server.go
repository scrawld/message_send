/* ######################################################################
# Author: (zhengfei@fcadx.cn)
# Created Time: 2018-11-23 19:33:58
# File Name: passport_server.go
# Description:
####################################################################### */

package passport_server

import (
	"fmt"
	"os"
	"time"

	"message_sender/libs/config"

	"github.com/ant-libs-go/pool"
	"github.com/satori/go.uuid"
	"github.com/smallnest/rpcx/client"
	"gitlab.com/feichi/fcad_thrift/libs/go/common"
)

func NewPool() *pool.Pool {
	return &pool.Pool{
		New: func(key string) interface{} {
			consul := []string{fmt.Sprintf("%s:%s", config.Get().Consul.Host, config.Get().Consul.Port)}
			d := client.NewConsulDiscovery(config.Get().Rpc["passport_server"].Node, "Server", consul, nil)
			cli := client.NewXClient("Server", client.Failover, client.RandomSelect, d, client.DefaultOption)
			return cli
		}}
}

var (
	Default *pool.Pool = NewPool()
	host, _            = os.Hostname()
)

func BuildCommonHeader() *common.Header {
	return &common.Header{
		Requester: fmt.Sprintf("message_sender#%s", host),
		Sessid:    uuid.NewV4().String(),
		Timestamp: time.Now().Unix(),
		Version:   100,
		Metadata:  map[string]string{}}
}

func BuildCommonHeaderWithToken(token string) *common.Header {
	return &common.Header{
		Requester: fmt.Sprintf("message_sender#%s", host),
		Sessid:    uuid.NewV4().String(),
		Timestamp: time.Now().Unix(),
		Version:   100,
		Metadata:  map[string]string{"token": token}}
}
