package utils

import (
	"crypto/tls"
	"fmt"
	"gf-vue-admin/server/library/global"
	"net/smtp"
	"strings"

	"github.com/jordan-wright/email"
)

var Email = new(_email)

type _email struct {
	_email *email.Email
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 发送测试邮件
func (e *_email) Test(subject string, body string) error {
	to := strings.Split(global.Config.Email.To, ",")
	return e.send(to, subject, body)
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 发送邮件
func (e *_email) ErrorToEmail(subject string, body string) error {
	to := strings.Split(global.Config.Email.To, ",")
	if to[len(to)-1] == "" { // 判断切片的最后一个元素是否为空,为空则移除
		to = to[:len(to)-1]
	}
	return e.send(to, subject, body)
}

func (e *_email) send(to []string, subject string, body string) error {
	auth := smtp.PlainAuth("", global.Config.Email.From, global.Config.Email.Secret, global.Config.Email.Host)
	e._email = email.NewEmail()
	e._email.From = global.Config.Email.GetFrom()
	e._email.To = to
	e._email.Subject = subject
	e._email.HTML = []byte(body)
	hostAddr := fmt.Sprintf("%s:%d", global.Config.Email.Host, global.Config.Email.Port)
	if global.Config.Email.IsSsl {
		return e._email.SendWithTLS(hostAddr, auth, &tls.Config{ServerName: global.Config.Email.Host})
	} else {
		return e._email.Send(hostAddr, auth)
	}
}
