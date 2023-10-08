package services

import (
	"fmt"

	constants "github.com/yornifpaz/back_noteapp/app/constant"
	"github.com/yornifpaz/back_noteapp/lib"
	"github.com/yornifpaz/back_noteapp/templates"
)

type IEmailService interface {
	SendEmail(email string, typeEmail constants.MIMEType, templateType constants.EmailTypesTemplates, data any) (message string, err error)
}

type EmailService struct {
	emailLibrary lib.IEmailLibrary
	templates    templates.IEmailTemplate
}

func (emailService *EmailService) SendEmail(email string, typeEmail constants.MIMEType, templateType constants.EmailTypesTemplates, data any) (message string, err error) {

	body, errTemplate := emailService.templates.GetEmailTemplate(templateType, data)
	if errTemplate != nil {

		return "", fmt.Errorf("error al obtener el template de %s: %w", templateType.GetSubject(), errTemplate)
	}

	message, err = emailService.sendEmail(typeEmail, email, templateType.GetSubject(), body)
	if err != nil {
		return "", fmt.Errorf("no se pudo enviar el correo de %s: %w", templateType.GetSubject(), err)
	}

	return message, nil
}

func (emailService *EmailService) sendEmail(typeEmail constants.MIMEType, email string, subject string, body string) (string, error) {

	configMsg := lib.EmailMessage{
		To:       email,
		Subject:  subject,
		Body:     body,
		TypeBody: typeEmail,
	}

	message, err := emailService.emailLibrary.Send(configMsg)

	return message, err
}

func NewEmailService(emailLibrary lib.IEmailLibrary, templates templates.IEmailTemplate) IEmailService {
	return &EmailService{emailLibrary: emailLibrary, templates: templates}
}
