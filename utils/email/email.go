package email

import (
	"gopkg.in/gomail.v2"
	"strconv"
)

func SendToMail(mailTo []string, subject, body string) error {
	// 定义邮箱服务器连接信息
	mailConn := map[string]string {
		"user": "han.yang@yottacloud.cn",
		"pass": "29VGYav7uXQQgkpa",
		"host": "smtp.exmail.qq.com",
		"port": "465",
	}
	port, _ := strconv.Atoi(mailConn["port"])
	m := gomail.NewMessage()
	m.SetHeader("From", mailConn["user"])
	m.SetHeader("To", mailTo...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := gomail.NewDialer(mailConn["host"], port,
		mailConn["user"], mailConn["pass"])
	err := d.DialAndSend(m)
	return err
}