package templates

import (
	"bytes"
	"html/template"
)

type IEmailTemplate interface {
	GetEmailTemplate(templateName string, data any) (template string, err error)
}
type EmailTemplate struct {
	template         *template.Template
	templateContents ITemplate
}

// GetEmailTemplate implements IEmailTemplate.
func (emailTemplate *EmailTemplate) GetEmailTemplate(templateName string, data any) (template string, err error) {

	newTemplate, err := emailTemplate.templateContents.GetTemplate(templateName)

	if err != nil {

		return "", err

	}

	tmpl, errTemplate := emailTemplate.template.New(templateName).Parse(newTemplate)

	if errTemplate != nil {

		return "", errTemplate
	}

	// Crear un buffer para almacenar la salida de la plantilla
	var resultBuffer bytes.Buffer

	// Aplicar los datos a la plantilla y escribir en el buffer
	errExecute := tmpl.Execute(&resultBuffer, data)

	if errExecute != nil {

		return "", errExecute
	}

	// Convertir el buffer a una cadena
	resultString := resultBuffer.String()

	return resultString, nil
}

func NewEmailTemplate() IEmailTemplate {
	var templateObject = template.New("emailTemplate")
	var templateContents = NewTemplate()
	return &EmailTemplate{templateContents: templateContents, template: templateObject}
}
