package latrappemelder

import (
	"github.com/pkg/errors"
	"gopkg.in/mail.v2"
)

// SendMail sends a mail.
func (m *LaTrappeMelder) SendMail(to string, subject string, body string) error {

	// Construct mail message
	msg := mail.NewMessage()
	msg.SetHeader("From", m.config.SMTP.FromEmail /*, m.config.Global.SMTP.FromName */)
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", body)

	d := mail.NewDialer(
		m.config.SMTP.Host,
		m.config.SMTP.Port,
		m.config.SMTP.User,
		m.config.SMTP.Password,
	)

	if m.config.SMTP.DisableTLS {
		d.StartTLSPolicy = mail.NoStartTLS
	}

	// Send the actual mail
	if err := d.DialAndSend(msg); err != nil {
		return errors.Wrap(err, "couldn't send the email")
	}

	return nil
}
