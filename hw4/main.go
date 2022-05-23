package main

import "workshop/hw4/email"

func main() {
	var s email.Email
	s = &email.Smtp{}
	s.Send([]string{}, "", "", "", "")
	s = &email.AwsSES{}
	s.Send([]string{}, "", "", "", "")
}
