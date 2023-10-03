package lib

type IEmailLibrary interface {
	Send(email string, subject string, body string) error
}

type EmailLibrary struct {
}

// send implements IEmailLibrary.
func (*EmailLibrary) Send(email string, subject string, body string) error {
	panic("unimplemented")
}

func NewEmailLibrary() IEmailLibrary {
	return &EmailLibrary{}
}
