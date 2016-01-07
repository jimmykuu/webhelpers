package webhelpers

import (
	"fmt"
	"net/smtp"
	"os/exec"
	"strings"
)

type SmtpConfig struct {
	Username string
	Password string
	Host     string
	Addr     string
}

// send mail
func SendMail(subject string, message string, from string, to []string, smtpConfig SmtpConfig, isHtml bool) error {
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
	msg := fmt.Sprintf("To: %s\r\nFrom: %s\r\nSubject: %s\r\nContent-Type: %s; charset=UTF-8\r\n\r\n%s", strings.Join(to, ";"), from, subject, contentType, message)
	return smtp.SendMail(smtpConfig.Addr, auth, from, to, []byte(msg))
}

// exec /usr/sbin/sendmail -t -i
func SendMailExec(subject string, message string, from string, to []string, sendmailPath string, isHtml bool) error {
	cmdArgs := strings.Fields(sendmailPath)
	cmdArgs = append(cmdArgs, "-f", from)
	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
	cmdStdin, err := cmd.StdinPipe()
	if err != nil {
		return err
	}
	if err := cmd.Start(); err != nil {
		return err
	}
	contentType := "text/plain"
	if isHtml {
		contentType = "text/html"
	}
	_, err = fmt.Fprintf(cmdStdin, "To: %s\r\nFrom: %s\r\nSubject: %s\r\nContent-Type: %s; charset=UTF-8\r\n\r\n%s", strings.Join(to, ";"), from, subject, contentType, message)
	err = cmdStdin.Close()
	if err != nil {
		return err
	}
	return cmd.Wait()
}
