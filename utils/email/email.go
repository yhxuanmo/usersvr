package email

import (
	"gopkg.in/gomail.v2"
	"strconv"
	"usersvr/utils"
)

func SendToMail(mailTo []string, subject, body string) error {
	// 定义邮箱服务器连接信息
	mailConn := map[string]string {
		"user": utils.Config.Email.User,
		"pass": utils.Config.Email.Pass,
		"host": utils.Config.Email.Host,
		"port": utils.Config.Email.Port,
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