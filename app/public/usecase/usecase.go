package usecase

import (
	"context"

	"github.com/bufbuild/protovalidate-go"
	"github.com/darwishdev/devkit-api-base/app/public/adapter"
	"github.com/darwishdev/devkit-api-base/app/public/repo"
	db "github.com/darwishdev/devkit-api-base/common/db/gen"
	abcv1 "github.com/darwishdev/devkit-api-base/common/pb/abc/v1"
	"github.com/darwishdev/devkit-api-base/common/supaclient"
)

type PublicUsecaseInterface interface {
	SettingsUpdate(ctx context.Context, req *abcv1.SettingsUpdateRequest) error
	SettingsFindForUpdate(ctx context.Context, req *abcv1.SettingsFindForUpdateRequest) (*abcv1.SettingsFindForUpdateResponse, error)
	IconsInputList(ctx context.Context) (*abcv1.IconsInputListResponse, error)
	UploadFile(ctx context.Context, req *abcv1.UploadFileRequest) (*abcv1.UploadFileResponse, error)
}

type PublicUsecase struct {
	repo       repo.PublicRepoInterface
	validator  *protovalidate.Validator
	supaClient supaclient.SupabaseServiceInterface
	adapter    adapter.PublicAdapterInterface
}

func NewPublicUsecase(store db.Store, validator *protovalidate.Validator, supaClient supaclient.SupabaseServiceInterface) *PublicUsecase {
	repo := repo.NewPublicRepo(store)
	adapter := adapter.NewPublicAdapter()

	return &PublicUsecase{
		repo:       repo,
		validator:  validator,
		supaClient: supaClient,
		adapter:    adapter,
	}
}
