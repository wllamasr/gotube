package mail

import (
	"net/smtp"
	"strings"
)

func SendMail(to, from, subject, body string) error {
	err := smtp.SendMail("stmp://localhost:1025",
		smtp.PlainAuth("", from, "", "localhost:1025"),
		from, strings.Fields(to), []byte(body))

	if err != nil {
		return err
	}
	return nil
}
