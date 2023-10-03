package template

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
	tmpl, err := emailTemplate.template.New(templateName).Parse(newTemplate)
	if err != nil {
		return "", err
	}

	// Crear un buffer para almacenar la salida de la plantilla
	var resultBuffer bytes.Buffer

	// Aplicar los datos a la plantilla y escribir en el buffer
	err = tmpl.Execute(&resultBuffer, data)
	if err != nil {
		return "", err
	}

	// Convertir el buffer a una cadena
	resultString := resultBuffer.String()
	return resultString, nil
}

func NewEmailTemplate(templateContents ITemplate, template *template.Template) IEmailTemplate {
	return &EmailTemplate{templateContents: templateContents, template: template}
}
