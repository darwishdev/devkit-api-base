package repo

import (
	"context"

	db "github.com/darwishdev/devkit-api-base/common/db/gen"
)

func (repo *AccountsRepo) OwnerFindByEmailOrCode(ctx context.Context, req string) (*db.AccountsSchemaOwnersView, error) {
	resp, err := repo.store.OwnerFindByEmailOrCode(context.Background(), req)
	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return &resp, nil
}

func (repo *AccountsRepo) OwnerFind(ctx context.Context, req int32) (*db.AccountsSchemaOwnersView, error) {
	resp, err := repo.store.OwnerFind(context.Background(), req)

	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return &resp, nil
}

func (repo *AccountsRepo) OwnersList(ctx context.Context) (*[]db.AccountsSchemaOwner, error) {
	resp, err := repo.store.OwnersList(context.Background())

	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return &resp, nil
}

func (repo *AccountsRepo) OwnersInputList(ctx context.Context) (*[]db.OwnersInputListRow, error) {
	resp, err := repo.store.OwnersInputList(context.Background())

	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return &resp, nil
}
func (repo *AccountsRepo) OwnerResetPassword(ctx context.Context, req *db.OwnerResetPasswordParams) error {
	err := repo.store.OwnerResetPassword(context.Background(), *req)

	if err != nil {
		return repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return nil
}

func (repo *AccountsRepo) OwnerDeleteRestore(ctx context.Context, req []int32) error {
	err := repo.store.OwnerDeleteRestore(context.Background(), req)
	if err != nil {
		return repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return nil
}

func (repo *AccountsRepo) OwnerCreate(ctx context.Context, req *db.OwnerCreateParams) (*db.AccountsSchemaOwner, error) {
	resp, err := repo.store.OwnerCreate(context.Background(), *req)

	if err != nil {

		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return &resp, nil
}

func (repo *AccountsRepo) OwnerFindForUpdate(ctx context.Context, req *int32) (*db.OwnerFindForUpdateRow, error) {
	resp, err := repo.store.OwnerFindForUpdate(context.Background(), *req)
	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return &resp, nil
}
func (repo *AccountsRepo) OwnerUpdate(ctx context.Context, req *db.OwnerUpdateParams) (*db.AccountsSchemaOwner, error) {
	resp, err := repo.store.OwnerUpdate(context.Background(), *req)
	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return &resp, nil
}
