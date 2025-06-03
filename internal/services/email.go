package services

import (
	"fmt"

	"go-fiber-boilerplate/internal/config"

	"gopkg.in/gomail.v2"
)

type EmailService struct {
	config *config.Config
}

func NewEmailService(cfg *config.Config) *EmailService {
	return &EmailService{config: cfg}
}

func (e *EmailService) SendActivationEmail(email, name, token string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", e.config.SMTPUsername)
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Account Activation")

	activationURL := fmt.Sprintf("%s/account_activations/%s/edit?email=%s",
		e.config.AppURL, token, email)

	body := fmt.Sprintf(`
		<h1>Sample App</h1>
		<p>Hi %s,</p>
		<p>Welcome to the Sample App! Click on the link below to activate your account:</p>
		<a href="%s">Activate</a>
		<p>If you have any questions, feel free to contact us.</p>
	`, name, activationURL)

	m.SetBody("text/html", body)

	d := gomail.NewDialer(e.config.SMTPHost, 587, e.config.SMTPUsername, e.config.SMTPPassword)
	return d.DialAndSend(m)
}

func (e *EmailService) SendPasswordResetEmail(email, name, token string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", e.config.SMTPUsername)
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Password Reset")

	resetURL := fmt.Sprintf("%s/password_resets/%s/edit?email=%s",
		e.config.AppURL, token, email)

	body := fmt.Sprintf(`
		<h1>Password Reset</h1>
		<p>Hi %s,</p>
		<p>To reset your password click the link below:</p>
		<a href="%s">Reset password</a>
		<p>This link will expire in two hours.</p>
		<p>If you did not request your password to be reset, please ignore this email and your password will stay as it is.</p>
	`, name, resetURL)

	m.SetBody("text/html", body)

	d := gomail.NewDialer(e.config.SMTPHost, 587, e.config.SMTPUsername, e.config.SMTPPassword)
	return d.DialAndSend(m)
}
