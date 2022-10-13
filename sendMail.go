package utils

import (
	"strconv"

	"gopkg.in/gomail.v2"
)

// Sender info.
type MailSender struct {
	Address  string // Sender email address.
	Nickname string // Define sender nickname.
	Password string // Mail password or token.
	Host     string // smtp.xxx.com
	Port     string // 465
}

type MailInfo struct {
	Receivers []string // receivers
	Subject   string
	Body      string
}

func SendMail(sender MailSender, info MailInfo) error {
	port, _ := strconv.Atoi(sender.Port)
	m := gomail.NewMessage()

	m.SetHeader("From", m.FormatAddress(sender.Address, sender.Nickname))
	m.SetHeader("To", info.Receivers...)
	m.SetHeader("Subject", info.Subject)
	m.SetBody("text/html", info.Body)

	d := gomail.NewDialer(sender.Host, port, sender.Address, sender.Password)
	err := d.DialAndSend(m)
	return err
}
