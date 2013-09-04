package webhelpers

import (
	"fmt"
	"net/smtp"
	"strings"
)

type SmtpConfig struct {
	Username string
	Password string
	Host     string
	Addr     string
}

// send mail
func SendMail(subject string, message string, from string, to []string, smtpConfig SmtpConfig, isHtml bool) {
	auth := smtp.PlainAuth(
		"",
		smtpConfig.Username,
		smtpConfig.Password,
		smtpConfig.Host,
	)
	contentType := "text/plain"
	if isHtml {
		contentType = "text/html"
	}
	msg := fmt.Sprintf("To: %s\r\nFrom: %s\r\nSubject: %s\r\nContent-Type: %s\r\n\r\n%s", strings.Join(to, ";"), from, subject, contentType, message)
	err := smtp.SendMail(smtpConfig.Addr, auth, from, to, []byte(msg))

	if err != nil {
		panic(err)
	}
}
