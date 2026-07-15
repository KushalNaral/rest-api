package tasks

import (
	"net/smtp"
	"strings"
	"testing"

	"rest-api/internal/config"
)

func TestSMTPEmailService_SendVerificationMail_MissingConfig(t *testing.T) {
	cfg := &config.Config{
		SMTPHost: "", // Missing host
	}

	service := NewSMTPEmailService(cfg)
	err := service.SendVerificationMail("test@example.com", "dummy-token")
	
	if err == nil {
		t.Fatal("Expected error due to missing SMTPHost, got nil")
	}

	if !strings.Contains(err.Error(), "SMTP configuration is missing") {
		t.Fatalf("Expected 'SMTP configuration is missing' error, got %v", err)
	}
}

func TestSMTPEmailService_HTMLBody(t *testing.T) {
	cfg := &config.Config{
		SMTPHost:        "localhost",
		SMTPPort:        "2525",
		SMTPFrom:        "noreply@example.com",
		CorsAllowOrigin: "http://localhost:5173",
	}

	// Mock the SendMailFunc
	var sentAddr, sentFrom string
	var sentTo []string
	var sentMsg []byte
	SendMailFunc = func(addr string, a smtp.Auth, from string, to []string, msg []byte) error {
		sentAddr = addr
		sentFrom = from
		sentTo = to
		sentMsg = msg
		return nil
	}
	defer func() {
		// Restore after test
		SendMailFunc = smtp.SendMail
	}()

	service := NewSMTPEmailService(cfg)
	err := service.SendVerificationMail("user@example.com", "secret-token")

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if sentAddr != "localhost:2525" {
		t.Errorf("Expected addr 'localhost:2525', got %v", sentAddr)
	}

	if sentFrom != "noreply@example.com" {
		t.Errorf("Expected from 'noreply@example.com', got %v", sentFrom)
	}

	if len(sentTo) != 1 || sentTo[0] != "user@example.com" {
		t.Errorf("Expected to 'user@example.com', got %v", sentTo)
	}

	msgStr := string(sentMsg)
	if !strings.Contains(msgStr, "http://localhost:5173/verify-email?token=secret-token") {
		t.Errorf("Expected message to contain verification link, got %v", msgStr)
	}
}
