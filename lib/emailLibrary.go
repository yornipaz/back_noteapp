package lib

import (
	"crypto/tls"
	"fmt"

	constants "github.com/yornifpaz/back_noteapp/app/constant"
	gomail "gopkg.in/gomail.v2"
)

// IEmailLibrary define la interfaz para la biblioteca de correo electrónico.
type IEmailLibrary interface {
	Send(message EmailMessage) (messageResponse string, err error)
}

type EmailLibrary struct {
	config EmailConfig
}
type EmailConfig struct {
	SMTPHost    string
	SMTPPort    int
	Username    string
	Password    string
	FromAddress string
	FromName    string
}
type EmailMessage struct {
	To       string
	Subject  string
	Body     string
	TypeBody constants.MIMEType
}

// newMessage implements IEmailLibrary.
func (emailLibrary *EmailLibrary) newMessage(emailMessage EmailMessage) (message *gomail.Message) {
	m := gomail.NewMessage()
	m.SetHeader("From", emailLibrary.config.FromAddress)
	m.SetHeader("To", emailMessage.To)
	m.SetHeader("Subject", emailMessage.Subject)
	m.SetBody(emailMessage.TypeBody.GetMIMEType(), emailMessage.Body)
	return m
}

// ConfigDialer implements IEmailLibrary.
func (*EmailLibrary) configDialer(config EmailConfig) *gomail.Dialer {
	d := gomail.NewDialer(config.SMTPHost, config.SMTPPort, config.Username, config.Password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	return d
}

// send implements IEmailLibrary.
func (emailLibrary *EmailLibrary) Send(message EmailMessage) (messageResponse string, err error) {
	dialer := emailLibrary.configDialer(emailLibrary.config)
	messages := emailLibrary.newMessage(message)
	return emailLibrary.send(dialer, messages)

}
func (*EmailLibrary) send(dialer *gomail.Dialer, messages ...*gomail.Message) (messageResponse string, err error) {
	err = dialer.DialAndSend(messages...)
	if err != nil {
		fmt.Println(" : " + err.Error())
		return "Error al enviar el correo : " + err.Error(), err
	}
	return "Correo enviado correctamente", nil

}
func NewEmailLibrary(config EmailConfig) IEmailLibrary {
	return &EmailLibrary{config: config}
}
