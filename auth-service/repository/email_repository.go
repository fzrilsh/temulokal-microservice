package repository

import (
	"fmt"
	"net/smtp"
)

type EmailData struct {
	To      string
	Subject string
	Body    string
}

type EmailRepository interface {
	SendEmail(data EmailData) error
}

type emailRepository struct {
	auth    *smtp.Auth
	address *string
	from    *string
}

// constructor
func NewEmailRepository(host string, port int, username string, password string, from string) EmailRepository {
	auth := smtp.PlainAuth("", username, password, host)
	address := fmt.Sprintf("%s:%d", host, port)
	return &emailRepository{
		auth:    &auth,
		address: &address,
		from:    &from,
	}
}

func (r *emailRepository) SendEmail(data EmailData) error {
	msg := []byte(fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n%s",
		data.To,
		data.Subject,
		data.Body,
	))

	if err := smtp.SendMail(*r.address, *r.auth, *r.from, []string{data.To}, msg); err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}
