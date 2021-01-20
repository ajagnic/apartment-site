package email

import (
	"net/smtp"
)

var (
	auth smtp.Auth
	addr string
	from string
)

// Initialize authentication values for the smtp client.
func Initialize(host, port, sender, password string) {
	addr = host + ":" + port
	from = sender
	auth = smtp.PlainAuth("", sender, password, host)
}

// Send content as an email to all recipients.
func Send(recipients []string, content []byte) error {
	err := smtp.SendMail(addr, auth, from, recipients, content)
	if err != nil {
		return err
	}
	return nil
}
