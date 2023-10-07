package lib

import (
	"crypto/tls"
	"fmt"

	gomail "gopkg.in/gomail.v2"
)

// IEmailLibrary define la interfaz para la biblioteca de correo electrónico.
type IEmailLibrary interface {
	Send(dialer *gomail.Dialer, messages ...*gomail.Message) (messageResponse string, err error)
	ConfigDialer() *gomail.Dialer
	CreateMessage(emailMessage EmailMessage) (message *gomail.Message)
}
type MIMEType string

const (
	// PlainText representa el tipo MIME para texto sin formato.
	PlainText MIMEType = "text/plain"

	// HTML representa el tipo MIME para contenido HTML.
	HTML MIMEType = "text/html"

	// OctetStream representa el tipo MIME para datos binarios sin formato.
	OctetStream MIMEType = "application/octet-stream"

	// PDF representa el tipo MIME para archivos PDF.
	PDF MIMEType = "application/pdf"

	// JPEG representa el tipo MIME para imágenes JPEG.
	JPEG MIMEType = "image/jpeg"

	// MultipartAlternative representa el tipo MIME para contenido multipart/alternative.
	MultipartAlternative MIMEType = "multipart/alternative"

	// MultipartMixed representa el tipo MIME para contenido multipart/mixed.
	MultipartMixed MIMEType = "multipart/mixed"
)

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
	TypeBody MIMEType
}

// CreateMessage implements IEmailLibrary.
func (emailLibrary *EmailLibrary) CreateMessage(emailMessage EmailMessage) (message *gomail.Message) {
	return emailLibrary.newMessage(emailMessage)
}

// newMessage implements IEmailLibrary.
func (emailLibrary *EmailLibrary) newMessage(emailMessage EmailMessage) (message *gomail.Message) {
	m := gomail.NewMessage()
	m.SetHeader("From", emailLibrary.config.FromAddress)
	m.SetHeader("To", emailMessage.To)
	m.SetHeader("Subject", emailMessage.Subject)
	m.SetBody(string(emailMessage.TypeBody), emailMessage.Body)
	return m
}

// ConfigDialer implements IEmailLibrary.
func (emailLibrary *EmailLibrary) ConfigDialer() *gomail.Dialer {
	return emailLibrary.configDialer(emailLibrary.config)
}

// ConfigDialer implements IEmailLibrary.
func (*EmailLibrary) configDialer(config EmailConfig) *gomail.Dialer {
	d := gomail.NewDialer(config.SMTPHost, config.SMTPPort, config.Username, config.Password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	return d
}

// send implements IEmailLibrary.
func (emailLibrary *EmailLibrary) Send(dialer *gomail.Dialer, messages ...*gomail.Message) (messageResponse string, err error) {
	return emailLibrary.send(dialer, messages...)

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
