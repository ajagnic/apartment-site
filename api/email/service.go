// Package email sends parsed HTML to recipients from the database.
package email

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/smtp"
	"os"
	"strconv"
	"time"

	"github.com/ajagnic/apartment-site/db"
)

var (
	host, port, from, pw, timeout, domain, addr string
	auth                                        smtp.Auth
)

func init() {
	host = os.Getenv("EMAIL_HOST")
	port = os.Getenv("EMAIL_PORT")
	from = os.Getenv("EMAIL_SENDER_ADDR")
	pw = os.Getenv("EMAIL_SENDER_PW")
	timeout = os.Getenv("EMAIL_PROCESS_TIME")
	domain = os.Getenv("SITE_DOMAIN")
	addr = host + ":" + port
	auth = smtp.PlainAuth("", from, pw, host)
}

// Process runs indefinitely, querying db records and emailing confirmations. (blocking)
func Process() {
	minutes, err := strconv.Atoi(timeout)
	if err != nil {
		log.Println(err)
		minutes = 5
	}
	for {
		time.Sleep(time.Duration(minutes) * time.Minute)
		confs, err := db.ReservationsToConfirm()
		if err != nil {
			log.Println(err)
		} else {
			for _, c := range confs {
				id := c.ID.Hex()
				err = sendConfirmation(id, c.Email)
				if err != nil {
					log.Printf("Error sending email for: %v:%v\n", id, c.Email)
				} else {
					db.SetBoolean(id, "emailed", true)
				}
			}
		}
	}
}

func sendConfirmation(id string, to string) (err error) {
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
	url := fmt.Sprintf("http://%s/confirmation?id=%s", domain, id)
	t.Execute(&body, struct{ URL string }{url})
	return body.Bytes(), nil
}
