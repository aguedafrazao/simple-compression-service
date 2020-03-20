package mail

import (
	"encoding/base64"
	"fmt"
	"net"
	"net/mail"
	"net/smtp"
	"strings"
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
	msg := composeMimeMail(to, c.Email, subject, body)
	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth(subject, c.Email, c.Password, "smtp.gmail.com"),
		c.Email, []string{to}, msg)
	if err != nil {
		fmt.Println("could not sent email: ", err)
	}
}
func getMXRecord(to string) (mx string, err error) {
	var e *mail.Address
	e, err = mail.ParseAddress(to)
	if err != nil {
		return
	}
	domain := strings.Split(e.Address, "@")[1]
	var mxs []*net.MX
	mxs, err = net.LookupMX(domain)
	if err != nil {
		return
	}
	for _, x := range mxs {
		mx = x.Host
		return
	}
	return
}

// Never fails, tries to format the address if possible
func formatEmailAddress(addr string) string {
	e, err := mail.ParseAddress(addr)
	if err != nil {
		return addr
	}
	return e.String()
}

func encodeRFC2047(str string) string {
	// use mail's rfc2047 to encode any string
	addr := mail.Address{Address: str}
	return strings.Trim(addr.String(), " <>")
}

func composeMimeMail(to string, from string, subject string, body string) []byte {
	header := make(map[string]string)
	header["From"] = formatEmailAddress(from)
	header["To"] = formatEmailAddress(to)
	header["Subject"] = encodeRFC2047(subject)
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/plain; charset=\"utf-8\""
	header["Content-Transfer-Encoding"] = "base64"
	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + base64.StdEncoding.EncodeToString([]byte(body))
	return []byte(message)
}
