package resend

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"

	db "github.com/darwishdev/devkit-api-base/common/db/gen"
	"github.com/resend/resend-go/v2"
	"github.com/rs/zerolog/log"
)

func (r *ResendService) SendEmail(req *resend.SendEmailRequest) (*resend.SendEmailResponse, error) {
	resp, err := r.Client.Emails.Send(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// LoadTemplate loads an email template from a file and replaces placeholders
func LoadTemplate(filePath string, data interface{}) (string, error) {
	tmplBytes, err := os.ReadFile(filepath.Clean(filePath))
	if err != nil {
		return "", err
	}

	tmpl, err := template.New("emailTemplate").Parse(string(tmplBytes))
	if err != nil {
		return "", err
	}

	var tplBuffer bytes.Buffer
	err = tmpl.Execute(&tplBuffer, data)
	if err != nil {
		return "", err
	}

	return tplBuffer.String(), nil
}
func (r *ResendService) SendReservationEventEmail(event string, reservationDetails db.ReservationsSchemaReservationsView) error {
	emailTemplates, ok := r.templates[event]
	if !ok {
		return fmt.Errorf("unknown event type: %s", event)
	}

	reservationDetails.PaymentUrl = fmt.Sprintf("%s%s", r.BaseUrl, reservationDetails.PaymentUrl)
	for _, template := range emailTemplates {
		htmlContent, err := LoadTemplate(template, reservationDetails)
		if err != nil {
			return err
		}

		to := []string{reservationDetails.CustomerEmail}
		isForOwner := strings.Contains(template, "owner")
		if isForOwner {
			to = []string{reservationDetails.OwnerEmail}

			if reservationDetails.RepresentativeOwnerEmail != "" {
				to = append(to, reservationDetails.RepresentativeOwnerEmail)
			}
		}

		eventName := strings.ReplaceAll(event, "_", " ")
		emailRequest := &resend.SendEmailRequest{
			From:    "noreply@yallabeina.com",
			To:      to,
			Html:    htmlContent,
			Subject: fmt.Sprintf("%s notification", eventName),
		}

		_, err = r.SendEmail(emailRequest)
		log.Debug().Interface("SendEmail", err).Msg("reservationDetails")
		if err != nil {
			return err
		}
		log.Debug().Interface("reservationDetails", isForOwner).Msg("reservationDetails")
	}
	return nil
}

func (r *ResendService) SendPaymentEventEmail(event string, reservationDetails db.ReservationsSchemaPaymentsView) error {
	emailTemplates, ok := r.templates[event]
	if !ok {
		return fmt.Errorf("unknown event type: %s", event)
	}

	for _, template := range emailTemplates {
		htmlContent, err := LoadTemplate(template, reservationDetails)
		if err != nil {
			return err
		}

		to := []string{reservationDetails.CustomerEmail}
		isForOwner := strings.Contains(template, "owner")
		if isForOwner {
			to = []string{reservationDetails.OwnerEmail}

			if reservationDetails.RepresentativeOwnerEmail != "" {
				to = append(to, reservationDetails.RepresentativeOwnerEmail)
			}
		}
		eventName := strings.ReplaceAll(event, "_", " ")
		emailRequest := &resend.SendEmailRequest{
			From:    "noreply@yallabeina.com",
			To:      to,
			Html:    htmlContent,
			Subject: fmt.Sprintf("%s notification", eventName),
		}

		_, err = r.SendEmail(emailRequest)
		if err != nil {
			return err
		}
		log.Debug().Interface("reservationDetails", isForOwner).Msg("reservationDetails")
	}
	return nil
}
