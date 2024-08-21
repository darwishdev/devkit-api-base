package supaclient

import (
	"context"
	"fmt"
	"io"

	"github.com/darwishdev/devkit-api-base/common/supaapi"
	"github.com/rs/zerolog/log"
	// supa "github.com/darwishdev/supaapi-go"
)

func (s *SupabaseService) SignIn(c context.Context, req supaapi.UserCredentials) (user *supaapi.AuthenticatedDetails, err error) {
	user, err = s.Client.Auth.SignIn(c, req)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *SupabaseService) UserUpdate(c context.Context, userToken string, updateData map[string]interface{}) (user *supaapi.User, err error) {

	user, err = s.Client.Auth.UpdateUser(c, userToken, updateData)
	if err != nil {
		return nil, err
	}

	return user, nil
}
func (s *SupabaseService) SignUp(c context.Context, req supaapi.UserCredentials) (user *supaapi.AuthenticatedDetails, err error) {
	user, err = s.Client.Auth.SignUp(c, req)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *SupabaseService) UserUpdateById(ctx context.Context, userID string, params supaapi.AdminUserParams) (*supaapi.AdminUser, error) {

	user, err := s.Client.Admin.UpdateUser(ctx, userID, params)
	if err != nil {
		log.Debug().Interface("weerro from UserUpdateById on supaservice.go", err).Interface("userId", userID).Interface("params", params).Msg("UserUpdateById")
		return nil, err
	}

	return user, nil
}

type UploadRequest struct {
	BucketName string
	Path       string
	Reader     io.Reader
	FileType   string
}

func (s *SupabaseService) Upload(req UploadRequest) (supaapi.FileResponse, error) {
	return s.Client.Storage.From(req.BucketName).Upload(req.Path, req.Reader, req.FileType), nil

}

type UploadMultipleRequest struct {
	Files []UploadRequest
}

func (s *SupabaseService) UploadMultiple(req UploadMultipleRequest) ([]*supaapi.FileResponse, error) {
	if len(req.Files) == 0 {
		return nil, fmt.Errorf("no files passed")
	}
	responses := make([]*supaapi.FileResponse, len(req.Files))
	for fileIndex, file := range req.Files {
		response := s.Client.Storage.From(file.BucketName).Upload(file.Path, file.Reader, file.FileType)
		responses[fileIndex] = &response
	}
	return responses, nil

}
