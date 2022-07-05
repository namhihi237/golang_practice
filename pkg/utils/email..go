package utils

import (
	"net/smtp"
	"practice/config"
)

const (
	host = "smtp.gmail.com"
	addr = "smtp.gmail.com:587"
)

// note: password used for gmail is application password
func sendEmail(email string, subject string, body string) error {
	env, err := config.GetEnv()
	if err != nil {
		return err
	}
	// Set up authentication information.
	auth := smtp.PlainAuth("", env.Email, env.EmailPass, host)

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	to := []string{email}
	msg := prepareMessage(email, subject, body)

	return smtp.SendMail(addr, auth, env.Email, to, msg)
}

func prepareMessage(email string, subject string, body string) []byte {
	return []byte("To: " + email + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		body + "\r\n")
}

func SendEmailActiveAccount(email string) error {
	subject := "Active your account"

	// generate token
	token, err := GenerateToken(map[string]interface{}{
		"id":        int64(0),
		"user_name": "",
		"email":     email,
	})

	if err != nil {
		return err
	}

	body := "Please active your account by click this link: http://localhost:8000/auth/active?token=" + *token

	return sendEmail(email, subject, body)
}
