package main

import (
	"testing"
	"gopkg.in/gomail.v2"
)
/*
url截取
*/

func Test_Mail(t *testing.T) {
	mailHeader := map[string][]string{
		"From":    {"z_haiyang@163.com"},
		"To":      {"z_haiyang@163.com"},
		"Subject": {"Hello"},
	}

	m := gomail.NewMessage()
	m.SetHeaders(mailHeader)
	// m.SetAddressHeader("Cc", "shengang@pingcap.com", "shengang")
	m.SetBody("text/html", "Hello <b>Ocean</b> <hr>  <div > <font color='red'>服务已挂，请速去查看!</font></div>")
	// m.Attach("../../go.mod")

	d := gomail.NewDialer("smtp.163.com", 465, "z_haiyang@163.com", "UOFHSQOFVZESIIEX")

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}