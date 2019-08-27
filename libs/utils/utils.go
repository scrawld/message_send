/* ######################################################################
# Author: (zfly1207@126.com)
# Created Time: 2018-09-12 16:44:18
# File Name: utils.go
# Description:
####################################################################### */

package utils

import (
	"encoding/base64"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

/**
 * assign one struct to other struct
 * @param: origin
 * @param: target
 * @params: excludes ... the attribute name exclude assign
 */
func Assign(origin, target interface{}, excludes ...string) {
	val_origin := reflect.ValueOf(origin).Elem()
	val_target := reflect.ValueOf(target).Elem()

	for i := 0; i < val_origin.NumField(); i++ {
		if !val_target.FieldByName(val_origin.Type().Field(i).Name).IsValid() {
			continue
		}
		is_exclude := false
		for _, col := range excludes {
			if val_origin.Type().Field(i).Name == col {
				is_exclude = true
				break
			}
		}
		if is_exclude {
			continue
		}
		switch val_origin.Field(i).Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			val_target.FieldByName(val_origin.Type().Field(i).Name).SetInt(val_origin.Field(i).Int())
		case reflect.String:
			val_target.FieldByName(val_origin.Type().Field(i).Name).SetString(val_origin.Field(i).String())
		}
	}
}

func EncodeId(id int64) (r string) {
	r = base64.StdEncoding.EncodeToString([]byte(strconv.FormatInt(id+33554432, 32)))
	return
}

func DecodeId(id string) (r int64, err error) {
	bs, err := base64.StdEncoding.DecodeString(id)
	if err != nil {
		err = fmt.Errorf("%s base64Decode, %s", id, err)
		return
	}
	r, err = strconv.ParseInt(string(bs), 32, 64)
	if err != nil {
		err = fmt.Errorf("%s parse, %s", id, err)
		return
	}
	// 32 to 10, - 32768
	r -= 33554432
	return
}

func VersionStringToInt(inp string) int32 {
	var s []string
	for _, v := range strings.Split(inp, ".") {
		if len(v) == 1 {
			v = fmt.Sprintf("0%s", v)
		}
		s = append(s, v)
	}
	r, _ := strconv.ParseInt(strings.Join(s, ""), 10, 64)
	return int32(r)
}

func VersionIntToString(inp int32) string {
	var s []string
	for _, v := range []int32{10000, 100, 1} {
		broken := inp / v
		inp -= broken * v
		s = append(s, strconv.Itoa(int(broken)))
	}
	return strings.Join(s, ".")
}

func DateRange(s, e string) (r []string) {
	st, _ := time.Parse("20060102", s)
	et, _ := strconv.Atoi(e)
	for {
		t, _ := strconv.Atoi(st.Format("20060102"))
		if t > et {
			break
		}
		r = append(r, strconv.Itoa(t))
		st = st.AddDate(0, 0, +1)
	}
	return
}

func FormatToMoney(inp float64) string {
	return fmt.Sprintf("%.2f", inp/1000000.0)
}

func FormatToPercent(inp float64) string {
	return fmt.Sprintf("%.2f%%", inp)
}

func GetImg(fn string) string {
	if fn == "" {
		return fn
	}
	if strings.HasPrefix(fn, "http") {
		return fn
	}
	baseUrl := "https://st.fcadx.cn/"
	return baseUrl + fn
}
