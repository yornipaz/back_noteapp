package templates

import (
	"bytes"
	"html/template"

	constants "github.com/yornifpaz/back_noteapp/app/constant"
)

type IEmailTemplate interface {
	GetEmailTemplate(templateType constants.EmailTypesTemplates, data any) (template string, err error)
}
type EmailTemplate struct {
	template         *template.Template
	templateContents ITemplate
}

// GetEmailTemplate implements IEmailTemplate.
func (emailTemplate *EmailTemplate) GetEmailTemplate(templateType constants.EmailTypesTemplates, data any) (template string, err error) {

	newTemplate, err := emailTemplate.templateContents.GetTemplate(templateType)

	if err != nil {

		return "", err

	}

	tmpl, errTemplate := emailTemplate.template.New(templateType.GetEmailTemplate()).Parse(newTemplate)

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
