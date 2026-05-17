//go:build mailgun_integration

package mailgun

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMailgunNotifier_SendNotificationByEmail(t *testing.T) {
	// given
	err := godotenv.Load("testdata/.env")
	require.NoError(t, err)

	contact := Contact{
		FirstName: "Christian",
		LastName:  "Sperandio",
		Email:     "christian.sperandio@gmail.com",
	}

	apiKey := os.Getenv("MAILGUN_API_KEY")
	domain := os.Getenv("MAILGUN_DOMAIN")

	notifier := NewMailgunNotifier(apiKey, domain)

	htmlContent := `<!DOCTYPE html>
<html>
<head>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Test Email</title>
</head>
<body style="font-family: Arial, sans-serif; line-height: 1.6; color: #333; max-width: 600px; margin: 0 auto; padding: 20px;">
    <h2 style="color: #2c3e50;">Test Notification</h2>
    <p>This is a test email sent for verification purposes only.</p>
    <p>No action is required. You can safely disregard this message.</p>
    <p>If you received this unexpectedly, please ignore it.</p>
    <hr style="border: none; border-top: 1px solid #eee; margin: 20px 0;">
    <p style="font-size: 0.8em; color: #7f8c8d;">
        This is an automated test notification.
    </p>
</body>
</html>`

	subject := "Action requise: Confirmation de votre présence"

	// when
	err = notifier.SendEmail(t.Context(), contact, htmlContent, "Text version", subject)

	// then
	assert.NoError(t, err)
}
