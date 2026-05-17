package mailgun

import (
	"context"
	"fmt"

	mailgunlib "github.com/mailgun/mailgun-go/v5"
	"github.com/rs/zerolog/log"
)

// Notifier handles sending emails via Mailgun.
type Notifier struct {
	apiKey string
	domain string
}

const apiBase = "https://api.eu.mailgun.net"

// Contact represents a person to whom an email can be sent.
type Contact struct {
	// FirstName is the contact's first name.
	FirstName string
	// LastName is the contact's last name.
	LastName string
	// Email is the contact's email address.
	Email string
}

// FullName returns the contact's full name.
func (c Contact) FullName() string {
	return fmt.Sprintf("%s %s", c.FirstName, c.LastName)
}

// SendEmail sends an email to the specified contact using the Mailgun API.
func (m *Notifier) SendEmail(ctx context.Context, contact Contact, htmlEmailContent string, textEmailContent string, subject string) error {
	to := fmt.Sprintf("%s <%s>", contact.FullName(), contact.Email)
	msg := mailgunlib.NewMessage(m.domain,
		"Lien2Vies <noreplys@lien2vies.fr>",
		"[Lien2Vies] "+subject,
		textEmailContent,
		to,
	)
	msg.SetHTML(htmlEmailContent)

	mg := mailgunlib.NewMailgun(m.apiKey)
	_ = mg.SetAPIBase(apiBase)

	messageResponse, err := mg.Send(ctx, msg)
	if err != nil {
		return fmt.Errorf("sending notification email: %w", err)
	}

	log.Info().Any("response", messageResponse).Msgf("sending notification: email sent to %s", to)
	return nil
}

// NewMailgunNotifier creates a new instance of Notifier.
func NewMailgunNotifier(apiKey string, domain string) *Notifier {
	return &Notifier{
		apiKey: apiKey,
		domain: domain,
	}
}
