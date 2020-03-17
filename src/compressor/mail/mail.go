package mail

import (
	"fmt"
	"net/smtp"
)

// Client stores email and password
type Client struct {
	Email    string
	Password string
}

// New creates a new client
func New(email, password string) *Client {
	return &Client{
		Email:    email,
		Password: password,
	}
}

// Send sends an email
func (c *Client) Send(to, subject, body string) {
	msg := "From: " + c.Email + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n" +
		body
	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth(subject, c.Email, c.Password, "smtp.gmail.com"),
		c.Email, []string{to}, []byte(msg))
	if err != nil {
		fmt.Println("could not sent email: ", err)
	}
}
