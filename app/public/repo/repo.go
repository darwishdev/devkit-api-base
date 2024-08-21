package repo

import (
	"context"

	db "github.com/darwishdev/devkit-api-base/common/db/gen"
)

type PublicRepoInterface interface {
	SettingsUpdate(ctx context.Context, req *db.SettingsUpdateParams) error
	IconsInputList(ctx context.Context) (*[]db.Icon, error)
	SettingsFindForUpdate(ctx context.Context) (*[]db.SettingsFindForUpdateRow, error)
}

type PublicRepo struct {
	store        db.Store
	errorHandler map[string]string
}

func NewPublicRepo(store db.Store) *PublicRepo {
	errorHandler := map[string]string{}
	return &PublicRepo{
		store:        store,
		errorHandler: errorHandler,
	}
}
