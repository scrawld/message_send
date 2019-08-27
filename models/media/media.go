/* ######################################################################
# File Name: models/media/media.go
# Author: Rain
# Main: jiayd163@163.com
# Created Time: 2019-02-19 10:50:18
####################################################################### */
package media

import (
	"sync"
	"time"

	"message_sender/libs/passport_server"

	"github.com/ant-libs-go/util/logs"
	uuid "github.com/satori/go.uuid"
	types "gitlab.com/feichi/fcad_thrift/libs/go/fcmp_passport_types"

	services "gitlab.com/feichi/fcad_thrift/libs/go/fcmp_passport_services"
)

var lock sync.Mutex
var medias = map[int32]*types.Media{}

func Start() {
	go func() {
		for {
			ms, err := passport_server.SearchMediaByParams(&services.SearchMediaParams{})
			if err != nil {
				time.Sleep(time.Second)
				continue
			}
			var expiresTimes = []int32{}
			lock.Lock()
			for _, m := range ms {
				medias[m.Id] = m
				expiresTimes = append(expiresTimes, m.ExpiresTime)
			}
			lock.Unlock()
			sleepTime := getSleepTime(expiresTimes)
			log := logs.New(uuid.NewV4().String())
			log.Debugf("[media]sleep time: %d", sleepTime)
			time.Sleep(time.Second * sleepTime)
		}
	}()
}

func GetAccessTokenById(id int32) (r string) {
	lock.Lock()
	if m, ok := medias[id]; ok {
		r = m.AccessToken
	}
	lock.Unlock()
	return
}

func getSleepTime(expiresTimes []int32) (r time.Duration) {
	var min int32
	for _, v := range expiresTimes {
		if v == 0 {
			continue
		}
		if min != 0 && v > min {
			continue
		}
		min = v
	}
	poor := int64(min) - time.Now().Unix()
	if poor > 300 {
		r = 300
	} else {
		r = 1
	}
	return
}
