package test

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
	"testing"

	"github.com/jordan-wright/email"
)

func TestSendEmail(t *testing.T) {
	e := email.NewEmail()
	e.From = "novo<zwx_info@163.com>"
	e.To = []string{"weixiong_zhu@163.com"}
	//e.Bcc = []string{"test_bcc@example.com"}
	//e.Cc = []string{"test_cc@example.com"}
	e.Subject = "email test"
	e.Text = []byte("Text Body is, of course, supported!")
	e.HTML = []byte("your code is : <h1>HelloWorld</h1>")
	//err := e.Send("smtp.163.com:587", smtp.PlainAuth("", "zwx_info@163.com", "YKSKCZQCFIPWREFS", "smtp.163.com"))
	// 如果出现EOF异常 则关闭ssl重试
	// 跳过验证
	err := e.SendWithTLS("smtp.163.com:587",
		smtp.PlainAuth("", "zwx_info@163.com", "YKSKCZQCFIPWREFS", "smtp.163.com"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.163.com"})
	if err != nil {
		fmt.Println(err)
	}
}
