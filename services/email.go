package services

import (
	"fmt"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendMagicLink(address string, tokenString string) error {
	from := mail.NewEmail("FeedMe", "shanecurtis1705@gmail.com")
	subject := "Login to FeedMe"
	to := mail.NewEmail(address, address)
	plainTextContent := "Click the button below to log in."

	// Format the magic link with the token
	magicLink := fmt.Sprintf(os.Getenv("Email_Verify_Link")+"/api/auth/verify?token=%s", tokenString)

	// HTML content with inline CSS for the button, and using the magic link
	htmlContent := fmt.Sprintf(`
	<div style="font-family: Arial, sans-serif;">
		<p>Click the button below to log in:</p>
		<a href="%s" 
		   style="display: inline-block; padding: 10px 20px; background-color: black; color: white; text-decoration: none; border-radius: 5px; font-weight: bold;">
		   Log in
		</a>
	</div>
	`, magicLink)

	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("Email_key"))
	response, err := client.Send(message)
	if err != nil {
		return err
	}
	println(response.StatusCode)
	return nil
}
