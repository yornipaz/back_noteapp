package services

import (
	"fmt"

	"github.com/yornifpaz/back_noteapp/lib"
	"github.com/yornifpaz/back_noteapp/template"
)

type IEmailService interface {
	SendEmail(email string, templateName string, subject Subject, data any) (message string, err error)
}

type EmailService struct {
	emailLibrary lib.EmailLibrary
	templates    template.IEmailTemplate
}

type Subject string

const (
	Recovery Subject = "Recuperación de contraseña"
	Welcome  Subject = "Bienvenido a Guaticer@"
)

func (emailService *EmailService) SendEmail(email string, templateName string, subject Subject, data any) (message string, err error) {
	body, errTemplate := emailService.templates.GetEmailTemplate(templateName, data)
	if errTemplate != nil {
		return "", fmt.Errorf("error al obtener el template de %s: %w", subject, errTemplate)
	}

	message, err = emailService.sendEmail(lib.HTML, email, string(subject), body)
	if err != nil {
		return "", fmt.Errorf("no se pudo enviar el correo de %s: %w", subject, err)
	}

	return message, nil
}

func (emailService *EmailService) sendEmail(typeEmail lib.MIMEType, email string, subject string, body string) (string, error) {
	dialer := emailService.emailLibrary.ConfigDialer()
	configMsg := lib.EmailMessage{
		To:       email,
		Subject:  subject,
		Body:     body,
		TypeBody: typeEmail,
	}
	msg := emailService.emailLibrary.CreateMessage(configMsg)
	message, err := emailService.emailLibrary.Send(dialer, msg)
	return message, err
}

func NewEmailService(emailLibrary lib.EmailLibrary, templates template.IEmailTemplate) IEmailService {
	return &EmailService{emailLibrary: emailLibrary, templates: templates}
}
