package config

import (
	"os"
	"strconv"

	"github.com/yornifpaz/back_noteapp/app/models"
)

func getDefaultEmailConfig() (emailConfig models.EmailConfig) {
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))
	smtpUsername := os.Getenv("SMTP_USERNAME")
	smtpPassword := os.Getenv("SMTP_PASSWORD")
	fromAddress := os.Getenv("FROM_ADDRESS")
	fromName := os.Getenv("FROM_NAME")

	emailConfig = models.EmailConfig{
		SMTPHost:    smtpHost,
		SMTPPort:    smtpPort,
		Username:    smtpUsername,
		Password:    smtpPassword,
		FromAddress: fromAddress,
		FromName:    fromName,
	}
	return
}
