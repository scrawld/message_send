/* ######################################################################
# File Name: libs/easywechat/template.go
# Author: Rain
# Main: jiayd163@163.com
# Created Time: 2019-02-19 14:17:01
####################################################################### */
package easywechat

import (
	"encoding/json"
	"fmt"

	"message_sender/libs/utils"
)

const (
	templateSendUrl = "https://api.weixin.qq.com/cgi-bin/message/wxopen/template/send"
)

type resSendMessage struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

type resError struct {
	ResJson string
	IsRetry int
}

func SendTemplateMessage(accessToken, openid, template_id, form_id, page string, data map[string]interface{}) (r resError, err error) {
	uri := fmt.Sprintf("%s?access_token=%s", templateSendUrl, accessToken)
	var postData = map[string]interface{}{
		"touser":      openid,
		"template_id": template_id,
		"page":        page,
		"form_id":     form_id,
		"data":        data,
	}
	byteRes, err := utils.PostJSON(uri, postData)
	if err != nil {
		r.IsRetry = 1
		return
	}
	var res = resSendMessage{}
	err = json.Unmarshal(byteRes, &res)
	if err != nil {
		r.IsRetry = 1
		return
	}
	r.ResJson = string(byteRes)
	if res.Errcode != 0 {
		// 40037: template_id不正确 41028: form_id不正确，或者过期 41029: form_id已被使用
		// 41030: page不正确 45009:接口调用超过限额（目前默认每个帐号日调用限额为100万）
		// 42001: access_token 无效
		if res.Errcode == 41028 || res.Errcode == 41029 || res.Errcode == 42001 {
			r.IsRetry = 1
		}
		err = fmt.Errorf("wx error: %s", r.ResJson)
		return
	}
	return
}
