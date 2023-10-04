package models

type RecoveryEmailData struct {
	Username  string
	ResetLink string
}
type TemplateData struct {
	TypeName string
	Contend  string
}
type WelcomeEmailData struct {
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
	From     string
	Subject  string
	Body     string
	TypeBody string
}
