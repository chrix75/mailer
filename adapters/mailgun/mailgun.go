package mailgun

import (
	"context"
	"fmt"

	mailgunlib "github.com/mailgun/mailgun-go/v5"
	"github.com/rs/zerolog/log"
)

type Notifier struct {
	apiKey string
	domain string
}

const apiBase = "https://api.eu.mailgun.net"

type Contact struct {
	FirstName string
	LastName  string
	Email     string
}

func (c Contact) FullName() string {
	return fmt.Sprintf("%s %s", c.FirstName, c.LastName)
}

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

func NewMailgunNotifier(apiKey string, domain string) *Notifier {
	return &Notifier{
		apiKey: apiKey,
		domain: domain,
	}
}
