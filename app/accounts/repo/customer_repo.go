package repo

import (
	"context"

	db "github.com/darwishdev/devkit-api-base/common/db/gen"
)

func (repo *AccountsRepo) CustomerFind(ctx context.Context, req db.CustomerFindParams) (*db.CustomerFindRow, error) {
	resp, err := repo.store.CustomerFind(context.Background(), req)
	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return &resp, nil
}
func (repo *AccountsRepo) CustomersList(ctx context.Context) (*[]db.AccountsSchemaCustomer, error) {
	resp, err := repo.store.CustomersList(context.Background())

	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return &resp, nil
}
func (repo *AccountsRepo) CustomerResetPassword(ctx context.Context, req *db.CustomerResetPasswordParams) error {
	err := repo.store.CustomerResetPassword(context.Background(), *req)

	if err != nil {
		return repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return nil
}

func (repo *AccountsRepo) CustomerDeleteRestore(ctx context.Context, req []int32) error {
	err := repo.store.CustomerDeleteRestore(context.Background(), req)
	if err != nil {
		return repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return nil
}

func (repo *AccountsRepo) CustomerCreateUpdate(ctx context.Context, req *db.CustomerCreateUpdateParams) (*db.AccountsSchemaCustomer, error) {
	resp, err := repo.store.CustomerCreateUpdate(context.Background(), *req)

	if err != nil {

		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return &resp, nil
}
