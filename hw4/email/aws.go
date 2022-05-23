package email

import "fmt"

type AwsSES struct{}

func (s *AwsSES) Send(to []string, name, from, subject, body string) error {
	fmt.Println("Sending email via AWS SES...")
	return nil
}
