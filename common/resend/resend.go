package resend

import (
	"path/filepath"

	db "github.com/darwishdev/devkit-api-base/common/db/gen"
	"github.com/resend/resend-go/v2"
)

// su "github.com/darwishdev/supabase-go"

type ResendServiceInterface interface {
	SendEmail(req *resend.SendEmailRequest) (*resend.SendEmailResponse, error)
	SendReservationEventEmail(event string, reservationDetails db.ReservationsSchemaReservationsView) error
	SendPaymentEventEmail(event string, paymentDetails db.ReservationsSchemaPaymentsView) error
}

type ResendService struct {
	Client    *resend.Client
	BaseUrl   string
	templates map[string][]string
}

func NewResendService(apiKey string, baseUrl string) (ResendServiceInterface, error) {
	templatePath := "common/resend/email_templates"
	client := resend.NewClient(apiKey)
	templates := map[string][]string{
		"reservation_created":  {filepath.Join(templatePath, "reservation_created_customer.html"), filepath.Join(templatePath, "reservation_created_owner.html")},
		"reservation_approved": {filepath.Join(templatePath, "reservation_approved_customer.html"), filepath.Join(templatePath, "reservation_approved_owner.html")},
		"payment_created":      {filepath.Join(templatePath, "payment_created_customer.html"), filepath.Join(templatePath, "payment_created_owner.html")},
		"payment_approved":     {filepath.Join(templatePath, "payment_approved_customer.html"), filepath.Join(templatePath, "payment_approved_owner.html")},
	}
	return &ResendService{
		Client:    client,
		BaseUrl:   baseUrl,
		templates: templates,
	}, nil
}
