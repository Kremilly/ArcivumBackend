package services

import (
	"fmt"
	"os"

	"arcivum/configs"

	"github.com/resend/resend-go/v2"
)

func SendEmail(toEmail, subject, htmlContent, plain string) (bool, error) {
	client := resend.NewClient(os.Getenv("RESEND_API_KEY"))

	params := &resend.SendEmailRequest{
		From:    fmt.Sprintf("%s <no-reply@%s>", configs.ProductName, configs.DomainName),
		To:      []string{toEmail},
		Subject: subject,
		Html:    htmlContent,
		Text:    plain,
	}

	_, err := client.Emails.Send(params)
	if err != nil {
		return false, err
	}

	return true, nil
}
