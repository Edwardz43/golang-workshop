package email

import "fmt"

type Email struct {
	sender  string
	address string
	to      []string
	subject string
	body    string
}

type Option func(*Email)

func From(name, email string) Option {
	return func(e *Email) {
		e.sender = name
		e.address = email
	}
}

func To(receiver []string) Option {
	return func(e *Email) {
		e.to = receiver
	}
}

func Subject(s string) Option {
	return func(e *Email) {
		e.subject = s
	}
}

func Body(b string) Option {
	return func(e *Email) {
		e.body = b
	}
}

func (e *Email) Send() error {
	fmt.Printf("Sending email from : %s, %s\n", e.sender, e.address)
	fmt.Println("Sending email to:", e.to)
	fmt.Println("Subject:", e.subject)
	fmt.Println("Body:", e.body)
	return nil
}

func New(opts ...Option) *Email {
	e := &Email{}
	for _, opt := range opts {
		opt(e)
	}
	return e
}
