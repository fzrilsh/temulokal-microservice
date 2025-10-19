package repository

import (
	"sync"

	"gopkg.in/gomail.v2"
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
	dialer *gomail.Dialer
	sender string
}

// constructor
func NewEmailRepository(host string, port int, sender string, password string) EmailRepository {
	dialer := gomail.NewDialer(host, port, sender, password)
	dialer.SSL = true

	return &emailRepository{
		dialer: dialer,
		sender: sender,
	}
}

func (r *emailRepository) SendEmail(data EmailData) error {
	m := gomail.NewMessage()
	m.SetHeader("From", r.sender)
	m.SetHeader("To", data.To)
	m.SetHeader("Subject", data.Subject)
	m.SetBody("text/html", data.Body)

	// setup wait group
	var wg sync.WaitGroup
	wg.Add(1)

	// send email with fire and forget system but still make sure the program is running.
	go func() {
		defer wg.Done()
		go r.dialer.DialAndSend(m)
	}()

	return nil
}
