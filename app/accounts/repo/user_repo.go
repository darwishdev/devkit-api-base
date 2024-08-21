package repo

import (
	"context"

	db "github.com/darwishdev/devkit-api-base/common/db/gen"
)

func (repo *AccountsRepo) UserFind(ctx context.Context, req db.UserFindParams) (*db.AccountsSchemaUser, error) {
	resp, err := repo.store.UserFind(context.Background(), req)
	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return &resp, nil
}

func (repo *AccountsRepo) AuthUserIDFindByEmail(ctx context.Context, req string) (*string, error) {
	id, err := repo.store.AuthUserIDFindByEmail(context.Background(), req)
	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return &id, nil
}
func (repo *AccountsRepo) UserPermissionsList(ctx context.Context, req int32) (*[]db.UserPermissionsListRow, error) {
	resp, err := repo.store.UserPermissionsList(context.Background(), req)

	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return &resp, nil
}

func (repo *AccountsRepo) UsersList(ctx context.Context) (*[]db.UsersListRow, error) {
	resp, err := repo.store.UsersList(context.Background())

	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return &resp, nil
}
func (repo *AccountsRepo) UserResetPassword(ctx context.Context, req *db.UserResetPasswordParams) error {
	err := repo.store.UserResetPassword(context.Background(), *req)

	if err != nil {
		return repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return nil
}

func (repo *AccountsRepo) UserDeleteRestore(ctx context.Context, req []int32) error {
	err := repo.store.UserDeleteRestore(context.Background(), req)
	if err != nil {
		return repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return nil
}

func (repo *AccountsRepo) UserCreateUpdate(ctx context.Context, req *db.UserCreateUpdateParams) (*db.AccountsSchemaUser, error) {
	resp, err := repo.store.UserCreateUpdate(context.Background(), *req)

	if err != nil {

		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return &resp, nil
}
func (repo *AccountsRepo) UserFindForUpdate(ctx context.Context, req *int32) (*db.UserFindForUpdateRow, error) {
	resp, err := repo.store.UserFindForUpdate(context.Background(), *req)
	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return &resp, nil
}
func (repo *AccountsRepo) UserFindNavigationBars(ctx context.Context, req *int32) (*[]db.UserFindNavigationBarsRow, error) {
	resp, err := repo.store.UserFindNavigationBars(context.Background(), *req)
	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return &resp, nil
}

func (repo *AccountsRepo) UserRolesList(ctx context.Context, req *int32) (*[]db.UserRolesListRow, error) {
	resp, err := repo.store.UserRolesList(context.Background(), *req)
	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return &resp, nil
}
