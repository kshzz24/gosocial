package utils

import (
	"fmt"
	"os"

	"gopkg.in/gomail.v2"
)

func SendPasswordResetEmail(toEmail, resetToken string) error {
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	smtpUsername := os.Getenv("SMTP_USERNAME")
	smtpPassword := os.Getenv("SMTP_PASSWORD")
	smtpFrom := os.Getenv("SMTP_FROM")
	frontendURL := os.Getenv("FRONTEND_URL")
	resetLink := fmt.Sprintf("%s/reset-password?token=%s", frontendURL, resetToken)

	m := gomail.NewMessage()
	m.SetHeader("From", smtpFrom)
	m.SetHeader("To", toEmail)
	m.SetHeader("Subject", "Password Reset Request - GoSocial")
	htmlBody := fmt.Sprintf(`
		<html>
			<body>
				<h2>Password Reset Request</h2>
				<p>You requested to reset your password for your GoSocial account.</p>
				<p>Click the link below to reset your password:</p>
				<p><a href="%s">Reset Password</a></p>
				<p>Or copy and paste this link in your browser:</p>
				<p>%s</p>
				<p>This link will expire in 1 hour.</p>
				<p>If you didn't request this, please ignore this email.</p>
				<hr>
				<p><small>GoSocial - Reddit Clone</small></p>
			</body>
		</html>
	`, resetLink, resetLink)

	m.SetBody("text/html", htmlBody)

	port := 587
	if smtpPort != "" {
		fmt.Sscanf(smtpPort, "%d", &port)
	}

	d := gomail.NewDialer(smtpHost, port, smtpUsername, smtpPassword)
	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}
