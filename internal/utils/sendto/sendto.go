package sendto

import (
	"GolangBackendEcommerce/global"
	"bytes"
	"fmt"
	"net/smtp"
	"strings"
	"text/template"

	"go.uber.org/zap"
)

type EmailAddress struct {
	Address string `json:"address"`
	Name    string `json:"name"`
}

type Mail struct {
	From    EmailAddress `json:"from"`
	To      []string     `json:"to"`
	Subject string       `json:"subject"`
	Body    string       `json:"html_content"`
}

func BuildMessage(mail Mail) string {
	msg := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n"
	msg += fmt.Sprintf("From: %s\r\n", mail.From.Address)
	msg += fmt.Sprintf("To: %s\r\n", strings.Join(mail.To, ";"))
	msg += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
	msg += fmt.Sprintf("\r\n%s\r\n", mail.Body)

	return msg
}

func SendTextEmailOTP(to []string, from string, otp string) error {
	contentEmail := Mail{
		From: EmailAddress{
			Address: from,
			Name:    "Golang Ecommerce",
		},
		To:      to,
		Subject: "Your OTP Code",
		Body:    fmt.Sprintf("<h1>Your OTP code is: %s</h1>", otp),
	}

	messageMail := BuildMessage(contentEmail)

	// send smtp
	authentication := smtp.PlainAuth("", global.Config.SMTP.Username, global.Config.SMTP.Password, global.Config.SMTP.Host)

	err := smtp.SendMail(global.Config.SMTP.Host+fmt.Sprintf(":%d", global.Config.SMTP.Port), authentication, from, to, []byte(messageMail))

	if err != nil {
		global.Logger.Error("Failed to send email:", zap.Error(err))
		return err
	}

	return nil
}

func SendTemplateEmailOTP(to []string, from string, nameTemplate string, dataTemplate map[string]interface{}) error {
	htmlBody, err := getMailTemplate(nameTemplate, dataTemplate)
	if err != nil {
		global.Logger.Error("Failed to parse email template:", zap.Error(err))
		return err
	}

	return send(to, from, htmlBody)
}

func getMailTemplate(nameTemplate string, dataTemplate map[string]interface{}) (string, error) {
	htmlTemplate := new(bytes.Buffer)

	t := template.Must(
		template.New(nameTemplate).
			ParseFiles("templates-email/" + nameTemplate),
	)

	err := t.Execute(htmlTemplate, dataTemplate)
	if err != nil {
		return "", err
	}

	return htmlTemplate.String(), nil
}

func send(to []string, from string, htmlTemplate string) error {
	contentEmail := Mail{
		From: EmailAddress{
			Address: from,
			Name:    "Golang Ecommerce",
		},
		To:      to,
		Subject: "OTP Verification",
		Body:    htmlTemplate,
	}

	messageMail := BuildMessage(contentEmail)

	// send smtp
	authentication := smtp.PlainAuth("", global.Config.SMTP.Username, global.Config.SMTP.Password, global.Config.SMTP.Host)

	err := smtp.SendMail(global.Config.SMTP.Host+fmt.Sprintf(":%d", global.Config.SMTP.Port), authentication, from, to, []byte(messageMail))

	if err != nil {
		global.Logger.Error("Failed to send email:", zap.Error(err))
		return err
	}

	return nil
}
