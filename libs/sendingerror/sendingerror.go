/* ######################################################################
# File Name: libs/sendingerror/sendingerror.go
# Author: Rain
# Main: jiayd163@163.com
# Created Time: 2019-02-20 17:13:56
####################################################################### */
package sendingerror

import (
	"fcmp_message_sender/models"
	"fmt"
)

type SendingError struct {
	SendStatus models.SendStatus
	IsRetry    int
	Response   string
}

func BuildErr(isRetry int, f string, v ...interface{}) *SendingError {
	o := &SendingError{
		SendStatus: models.SendStatusFail,
		IsRetry:    isRetry,
		Response:   fmt.Sprintf(f, v...),
	}
	return o
}
