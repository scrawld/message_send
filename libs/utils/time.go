package utils

import (
	"strconv"
	"time"
)

// 获取今天的零点时间
func GetTodayBegin() int64 {
	timeStr := time.Now().Format("2006-01-02")
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr+" 00:00:00", time.Local)
	return t.Unix()
}

// 获取今天的最后结束时间
func GetTodayEnd() int64 {
	timeStr := time.Now().Format("2006-01-02")
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr+" 23:59:59", time.Local)
	return t.Unix()
}

func GetDt() int {
	dt, _ := strconv.Atoi(time.Now().Format("20060102"))
	return dt
}
