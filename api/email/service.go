package email

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
	"os"
)

var (
	host, port, from, pw, addr string
	linkURL                    string
	auth                       smtp.Auth
)

func init() {
	host = os.Getenv("EMAIL_HOST")
	port = os.Getenv("EMAIL_PORT")
	from = os.Getenv("EMAIL_SENDER_ADDR")
	pw = os.Getenv("EMAIL_SENDER_PW")
	domain := os.Getenv("SITE_DOMAIN")
	linkURL = fmt.Sprintf("http://%s/confirmation?id=", domain)
	addr = host + ":" + port
	auth = smtp.PlainAuth("", from, pw, host)
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
	url := linkURL + id
	t.Execute(&body, struct{ URL string }{url})
	return body.Bytes(), nil
}
