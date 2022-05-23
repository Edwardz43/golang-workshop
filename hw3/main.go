package main

import "workshop/hw3/email"

func main() {
	e := email.New(
		email.From("Ed Lo", "test@example.com"),
		email.To([]string{"John@example.co"}),
		email.Subject("Hello"),
		email.Body("Hello, World!"),
	)
	e.Send()
}
