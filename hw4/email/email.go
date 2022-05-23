package email

type Email interface {
	Send(to []string, name, from, subject, body string) error
}
