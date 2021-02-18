package sendmail

import (
	"GB/lesson-2/shop/models"
	"fmt"
	"net/smtp"
)

type EmailClient interface {
	SendOrderConfirmation(order *models.Order) error
}

type emailClient struct {
	cli *smtp.Client
}

// func Sendmail(email *emailClient) SendOrderConfirmation(order *models.Order) error {
//Sendmail sends e-mail
func Sendmail() {
	// Sender data.
	from := "s0ul.tr4d3r@gmail.com"
	password := "zfniagwdtvwxhqzc"

	// Receiver email address.
	to := []string{
		"a.bondar@corp.vk.com",
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Message.
	message := []byte("This is a test email message.")

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent Successfully!")
}
