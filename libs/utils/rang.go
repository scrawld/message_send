/* ######################################################################
# File Name: libs/utils/rang.go
# Author: Rain
# Main: jiayd163@163.com
# Created Time: 2019-01-25 18:50:39
####################################################################### */
package utils

import (
	"math/rand"
	"time"
)

/**
生成区间随机数
*/
func RandNum(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min+1)
}
