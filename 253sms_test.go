package sms253

import (
	"fmt"
	"testing"
)

func Test_two(t *testing.T) {

	sms := NewSms253("N00xxxxxx", "eLENxxxxxx")

	msg := "【253云通讯】" + "测试人员您的验ddd证码为888888请在5分钟内输入。"

	rl, err := sms.SendSms(msg, "13600xxxxx")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(rl)
}
