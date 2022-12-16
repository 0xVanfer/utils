package utils

import (
	"strconv"

	"gopkg.in/gomail.v2"
)

// Sender info.
type MailSender struct {
	Address  string `gorm:"column:address" json:"address"`   // Sender email address.
	Nickname string `gorm:"column:nickname" json:"nickname"` // Define sender nickname.
	Password string `gorm:"column:password" json:"password"` // Mail password or token.
	Host     string `gorm:"column:host" json:"host"`         // smtp.xxx.com
	Port     string `gorm:"column:port" json:"port"`         // 465
}

type MailInfo struct {
	Receivers []string // receivers
	Subject   string
	Body      string
}

// Deprecated. Will use new structure later.
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
