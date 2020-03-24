# Mail package

To use it import to your project:
```
$ go get github.com/ABuarque/simple-compression-service/src/libs/mail
```

It provides an API to send text emails using gmail account. To create a client use New functions:
```
// New creates a new client
func New(email, password string) *Client {
	return &Client{
		Email:    email,
		Password: password,
	}
}
```
To send email call Send:
```
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
```
