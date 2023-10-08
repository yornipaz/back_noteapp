package constants

type EmailTypesTemplates int
type MIMEType int

//constant Email Templates
const (
	Welcome EmailTypesTemplates = iota + 1
	ForgotPassword
)

const (
	// PlainText representa el tipo MIME para texto sin formato.
	PlainText MIMEType = iota + 1
	HTML
	OctetStream
	PDF
	JPEG
	MultipartAlternative
	MultipartMixed
)

func (e EmailTypesTemplates) GetEmailTemplate() string {

	return [...]string{"Welcome_user", "forgot_password"}[e-1]
}
func (e EmailTypesTemplates) GetSubject() string {
	return [...]string{"Bienvenido a Guaticer@", "Recuperación de contraseña"}[e-1]
}
func (m MIMEType) GetMIMEType() string {
	return [...]string{"text/plain", "text/html", "application/octet-stream", "application/pdf", "image/jpeg", "multipart/alternative", "multipart/mixed"}[m-1]

}
