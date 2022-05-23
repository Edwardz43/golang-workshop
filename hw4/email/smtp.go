package email

import "fmt"

type Smtp struct{}

func (s *Smtp) Send(to []string, name, from, subject, body string) error {
	// TODO
	fmt.Println("Sending email with SMTP...")
	return nil
}
