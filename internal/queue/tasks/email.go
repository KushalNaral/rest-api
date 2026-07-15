package tasks

import (
	"fmt"
	"net/smtp"
	"time"

	"rest-api/internal/config"
)

// SendMailFunc allows overriding smtp.SendMail for testing
var SendMailFunc = smtp.SendMail

type SMTPEmailService struct {
	cfg *config.Config
}

func NewSMTPEmailService(cfg *config.Config) *SMTPEmailService {
	return &SMTPEmailService{cfg: cfg}
}

// SendVerificationMail sends an email with the verification link to the user.
func (s *SMTPEmailService) SendVerificationMail(email, token string) error {
	if s.cfg.SMTPHost == "" || s.cfg.SMTPPort == "" {
		return fmt.Errorf("SMTP configuration is missing, cannot send email")
	}

	verificationLink := fmt.Sprintf("%s/verify-email?token=%s", s.cfg.CorsAllowOrigin, token)

	// Build proper RFC 5322 compliant message
	msg := s.buildVerificationEmail(email, verificationLink, token)

	auth := smtp.PlainAuth("", s.cfg.SMTPUser, s.cfg.SMTPPass, s.cfg.SMTPHost)
	addr := fmt.Sprintf("%s:%s", s.cfg.SMTPHost, s.cfg.SMTPPort)

	err := SendMailFunc(addr, auth, s.cfg.SMTPFrom, []string{email}, msg)
	if err != nil {
		return fmt.Errorf("error sending email: %w", err)
	}

	fmt.Printf("[email] verification email sent to %s\n", email)
	return nil
}

func (s *SMTPEmailService) buildVerificationEmail(to, verificationLink, token string) []byte {
	subject := "Verify your email address"

	body := fmt.Sprintf(`<html>
<head>
    <meta charset="UTF-8">
</head>
<body style="font-family: Arial, sans-serif; max-width: 600px; margin: 0 auto;">
    <h2>Welcome!</h2>
    <p>Please click the link below to verify your email address:</p>
    <p><a href="%s" style="color: #0066cc;">%s</a></p>
    <p>Or copy and paste this token manually: <strong>%s</strong></p>
    <hr>
    <p style="color: #666; font-size: 0.9em;">If you didn't request this verification, please ignore this email.</p>
</body>
</html>`, verificationLink, verificationLink, token)

	// Proper email headers
	headers := map[string]string{
		"From":         s.cfg.SMTPFrom,
		"To":           to,
		"Subject":      subject,
		"MIME-Version": "1.0",
		"Content-Type": "text/html; charset=UTF-8",
		"Date":         time.Now().UTC().Format(time.RFC1123Z),
	}

	var msg []byte
	for k, v := range headers {
		msg = append(msg, fmt.Sprintf("%s: %s\r\n", k, v)...)
	}
	msg = append(msg, "\r\n"...) // blank line between headers and body
	msg = append(msg, []byte(body)...)

	return msg
}
