package email

import (
	"context"
	"github.com/mailgun/mailgun-go/v3"
	"os"
	"time"
)

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
