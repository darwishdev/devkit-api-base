package supaclient

import (
	"context"

	"github.com/darwishdev/devkit-api-base/common/supaapi"
)

// su "github.com/darwishdev/supabase-go"

type SupabaseServiceInterface interface {
	SignUp(c context.Context, req supaapi.UserCredentials) (user *supaapi.AuthenticatedDetails, err error)
	SignIn(c context.Context, req supaapi.UserCredentials) (user *supaapi.AuthenticatedDetails, err error)
	UserUpdate(c context.Context, userToken string, updateData map[string]interface{}) (user *supaapi.User, err error)
	UserUpdateById(ctx context.Context, userID string, params supaapi.AdminUserParams) (*supaapi.AdminUser, error)
	Upload(req UploadRequest) (supaapi.FileResponse, error)
	UploadMultiple(req UploadMultipleRequest) ([]*supaapi.FileResponse, error)
	// Upload(ctx context.Context, file bytes) (string, error)
}

type SupabaseService struct {
	Client *supaapi.Client
}

func NewSupabaseService(supabaseUrl string, supabaseKey string) (SupabaseServiceInterface, error) {
	supa := supaapi.CreateClient(supabaseUrl, supabaseKey)

	return &SupabaseService{
		Client: supa,
	}, nil
}
