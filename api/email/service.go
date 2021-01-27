package email

import (
	"bytes"
	"html/template"
	"net/smtp"
)

var (
	auth smtp.Auth
	addr string
	from string
)

// TemplateData stores values to be injected into HTML.
type TemplateData struct {
	URL string
}

// Initialize authentication values for the smtp client.
func Initialize(host, port, sender, password string) {
	from = sender
	addr = host + ":" + port
	auth = smtp.PlainAuth("", sender, password, host)
}

// SendConfirmation emails a confirmation link to the recipient.
func SendConfirmation(id string, to string) (err error) {
	msg, err := generateEmailBody(id)
	if err != nil {
		return
	}
	recipients := []string{to}
	err = smtp.SendMail(addr, auth, from, recipients, msg)
	return
}

func generateEmailBody(id string) ([]byte, error) {
	var body bytes.Buffer
	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte("From: Agnic Apartments\r\nSubject: Confirm Your Reservation\r\n" + mimeHeaders))
	t, err := template.ParseFiles("./email/template.html")
	if err != nil {
		return nil, err
	}
	url := "http://localhost:3000/confirmation?id=" + id
	t.Execute(&body, TemplateData{url})
	return body.Bytes(), nil
}
