/* ######################################################################
# File Name: libs/easywechat/checked.go
# Author: Rain
# Main: jiayd163@163.com
# Created Time: 2019-01-07 10:40:06
####################################################################### */
package easywechat

import (
	"encoding/json"
	"fmt"
	"strings"

	"message_sender/libs/utils"
)

const (
	msgCheckUrl = "https://api.weixin.qq.com/wxa/msg_sec_check"
)

// Wechat struct
type Wechat struct {
	AccessToken string
}

func New(accessToken string) *Wechat {
	o := &Wechat{}
	o.AccessToken = accessToken
	return o
}

// reqMsgChecked 文本内容安全请求信息
type reqMsgChecked struct {
	Content string `json:"content"`
}

// resChecked 内容安全返回结果
type resChecked struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

func (this *Wechat) MsgChecked(content ...string) (err error) {
	req := &reqMsgChecked{strings.Join(content, "")}
	uri := fmt.Sprintf("%s?access_token=%s", msgCheckUrl, this.AccessToken)
	resp, err := utils.PostJSON(uri, req)
	if err != nil {
		return
	}
	var res resChecked
	if err = json.Unmarshal(resp, &res); err != nil {
		return
	}

	if res.Errcode != 0 {
		err = fmt.Errorf("MsgCheck error : errcode=%v , errmsg=%v", res.Errcode, res.Errmsg)
		return
	}

	return nil
}
