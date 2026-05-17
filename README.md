# Mailer Module

The `mailer` module provides email notification services for the application. It is designed to be an extensible component for sending transactional emails.

## Purpose

The primary purpose of this module is to abstract the complexities of email delivery and provide a simple interface to send emails to users.

## How it works

Currently, the module provides a **Mailgun** adapter. It uses the [Mailgun Go SDK](https://github.com/mailgun/mailgun-go) to communicate with the Mailgun API.

### Key Components

- **`mailgun.Notifier`**: The main struct that handles email delivery. It requires a Mailgun API key and a verified domain.
- **`mailgun.Contact`**: Represents a recipient with their first name, last name, and email address.
- **`SendEmail`**: The core method that sends an email with both HTML and plain text versions.

### Configuration

The Mailgun adapter is configured to use the **EU API base** (`https://api.eu.mailgun.net`).

## Usage

To use the mailer in your service, initialize the `MailgunNotifier` and call `SendEmail`:

```go
import "mailer/adapters/mailgun"

// Initialize the notifier
notifier := mailgun.NewMailgunNotifier(apiKey, domain)

// Define the recipient
contact := mailgun.Contact{
    FirstName: "John",
    LastName:  "Doe",
    Email:     "john.doe@example.com",
}

// Send an email
err := notifier.SendEmail(
    ctx,
    contact,
    "<html><body><h1>Hello!</h1></body></html>", // HTML content
    "Hello!",                                     // Text content
    "Subject of the email",
)
```
