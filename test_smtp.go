package main

import (
	"fmt"
	"net/smtp"
)

func main() {
	host := "sandbox.smtp.mailtrap.io"
	port := "2525"
	user := "82cd557299784e"
	pass := "b94338f7227507"

	auth := smtp.PlainAuth("", user, pass, host)
	addr := fmt.Sprintf("%s:%s", host, port)

	msg := []byte("To: test@example.com\r\n" +
		"Subject: Test\r\n" +
		"\r\n" +
		"This is a test email.\r\n")

	err := smtp.SendMail(addr, auth, "noreply@example.com", []string{"test@example.com"}, msg)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Println("Success!")
	}
}
