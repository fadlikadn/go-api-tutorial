package email

import (
	"context"
	"fmt"
	"github.com/mailgun/mailgun-go/v3"
	"net/smtp"
	"os"
	"time"
)

type smtpServer struct {
	host string
	port string
}

func (s *smtpServer) serverName() string {
	return s.host + ":" + s.port
}

func SendActivationEmail(receiver string) (string, error) {
	//mg := mailgun.NewMailgun(os.Getenv("MAILGUN_API_DOMAIN"), os.Getenv("MAILGUN_API_KEY"), os.Getenv("MAILGUN_PUBLIC_KEY"))
	mg := mailgun.NewMailgun(os.Getenv("MG_DOMAIN"), os.Getenv("MG_API_KEY"))
	//mg := mailgun.NewMailgun("https://api.mailgun.net/v3/sandbox3b93dfda06d3417b9c8d7c50141a5b30.mailgun.org", "key-f9f9f275cf228b5e7a82d4916432aaec")
	sender := os.Getenv("MAILGUN_SENDER")
	//sender := "info@billing-subscription.com"
	subject := "Activation Email"
	body := "Thank you for registration"
	recipient := receiver

	// The message object allows you to add attachment and Bcc recipients
	message := mg.NewMessage(sender, subject, body, recipient)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// Send the message with a 10 second timeout
	_, id, err := mg.Send(ctx, message)

	return id, err
}

func SendStatusEmail(receiver string) (int, error) {
	CONFIG_SMTP_HOST := os.Getenv("CONFIG_SMTP_HOST")
	CONFIG_SMTP_PORT := os.Getenv("CONFIG_SMTP_PORT")
	CONFIG_EMAIL := os.Getenv("CONFIG_EMAIL")
	CONFIG_EMAIL_PASSWORD := os.Getenv("CONFIG_EMAIL_PASSWORD")

	to := receiver
	from := os.Getenv("CONFIG_EMAIL")
	fmt.Println(receiver)
	fmt.Println(CONFIG_SMTP_HOST)
	fmt.Println(CONFIG_SMTP_PORT)
	fmt.Println(CONFIG_EMAIL)
	fmt.Println(CONFIG_EMAIL_PASSWORD)

	smtpServer := smtpServer{
		host: os.Getenv("CONFIG_SMTP_HOST"),
		port: os.Getenv("CONFIG_SMTP_PORT"),
	}

	message := []byte("Test sending email using gmail smtp")

	auth := smtp.PlainAuth("", os.Getenv("CONFIG_EMAIL"), os.Getenv("CONFIG_EMAIL_PASSWORD"), smtpServer.host)

	// sending email
	err := smtp.SendMail(smtpServer.serverName(), auth, from, []string{to}, message)
	/*err := smtp.SendMail(CONFIG_SMTP_HOST + ":" + CONFIG_SMTP_PORT,
		smtp.PlainAuth("", CONFIG_EMAIL, CONFIG_EMAIL_PASSWORD, CONFIG_SMTP_HOST),
		CONFIG_EMAIL,
		[]string{to},
		[]byte(message))*/
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	return 1, nil
}
