package services

import (
	"bytes"
	"html/template"
	"net/smtp"
	"os"

	"github.com/haji-sudo/ShabehRoshan/config"
)

func SendingVerificationEmail(to, username, token string) error {
	data := struct {
		Username         string
		VerificationLink string
	}{
		Username:         username,
		VerificationLink: "http://localhost:3000/verify-email?token=" + token,
	}
	// Render the email template with the dynamic data
	var renderedEmail bytes.Buffer

	if err := config.Engine.Render(&renderedEmail, "email/verify.email", data); err != nil {
		return err
	}
	// Get the final rendered email HTML
	body := renderedEmail.String()
	from := os.Getenv("GMAIL")
	pass := os.Getenv("GMAIL_APP_PASSWORD")
	headers := "From: " + from + "\r\n" +
		"To: " + to + "\r\n" +
		"Subject: Confirm Email\r\n" +
		"MIME-Version: 1.0\r\n" +
		"Content-Type: text/html; charset=\"UTF-8\"\r\n" +
		"\r\n"
	msg := headers + body
	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))
	if err != nil {
		return err
	}
	return nil
}

func SendingResetPasswordEmail(to, username, token string) error {
	templateFile := "./reset.password.html"
	templateBytes, err := os.ReadFile(templateFile)
	if err != nil {
		return err
	}
	templateString := string(templateBytes)
	tmpl, err := template.New("emailTemplate").Parse(templateString)
	if err != nil {
		return err
	}
	data := struct {
		Username  string
		ResetLink string
	}{
		Username:  username,
		ResetLink: "http://localhost:3000/forgot-password?token=" + token,
	}
	// Render the email template with the dynamic data
	var renderedEmail bytes.Buffer
	if err := tmpl.Execute(&renderedEmail, data); err != nil {
		return err
	}
	// Get the final rendered email HTML
	body := renderedEmail.String()
	from := "shabeh.roshan.email@gmail.com"
	pass := "ggzhktcnfjavikhf"
	headers := "From: " + from + "\r\n" +
		"To: " + to + "\r\n" +
		"Subject: Password Reset\r\n" +
		"MIME-Version: 1.0\r\n" +
		"Content-Type: text/html; charset=\"UTF-8\"\r\n" +
		"\r\n"
	msg := headers + body
	err = smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))
	if err != nil {
		return err
	}
	return nil
}
