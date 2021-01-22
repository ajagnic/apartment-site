package email

import (
	"fmt"
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

// SendConfirmation sends content as an email to all recipients.
func SendConfirmation(id string, to string) error {
	recipients := []string{to}
	body := fmt.Sprintf("Reservation ID: %s\r\n", id)
	msg := []byte("From: Agnic Apartments\r\n" +
		"Subject: Confirm Your Reservation\r\n" +
		"\r\n" + body)
	err := smtp.SendMail(addr, auth, from, recipients, msg)
	if err != nil {
		return err
	}
	return nil
}
