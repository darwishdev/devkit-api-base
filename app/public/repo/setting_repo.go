package repo

import (
	"context"

	db "github.com/darwishdev/devkit-api-base/common/db/gen"
)

func (repo *PublicRepo) SettingsUpdate(ctx context.Context, req *db.SettingsUpdateParams) error {
	_, err := repo.store.SettingsUpdate(context.Background(), *req)

	if err != nil {
		return repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return nil
}
func (repo *PublicRepo) SettingsFindForUpdate(ctx context.Context) (*[]db.SettingsFindForUpdateRow, error) {
	resp, err := repo.store.SettingsFindForUpdate(context.Background())
	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return &resp, nil
}

func (repo *PublicRepo) IconsInputList(ctx context.Context) (*[]db.Icon, error) {
	resp, err := repo.store.IconsInputList(context.Background())

	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}
	return &resp, nil
}
