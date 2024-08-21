package repo

import (
	"context"

	db "github.com/darwishdev/devkit-api-base/common/db/gen"
)

func (repo *AccountsRepo) PermissionsList(ctx context.Context) ([]db.PermissionsListRow, error) {
	resp, err := repo.store.PermissionsList(context.Background())
	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return resp, nil
}

func (repo *AccountsRepo) RolesInputList(ctx context.Context) (*[]db.RolesInputListRow, error) {
	resp, err := repo.store.RolesInputList(context.Background())
	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return &resp, nil
}

func (repo *AccountsRepo) RolesList(ctx context.Context) (*[]db.RolesListRow, error) {
	resp, err := repo.store.RolesList(context.Background())

	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return &resp, nil
}

func (repo *AccountsRepo) RoleCreate(ctx context.Context, req *db.RoleCreateTXParams) (*db.AccountsSchemaRole, error) {
	resp, err := repo.store.RoleCreateTX(context.Background(), *req)

	if err != nil {

		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return &resp.Role, nil
}

func (repo *AccountsRepo) RoleFindForUpdate(ctx context.Context, req *int32) (*[]byte, error) {
	resp, err := repo.store.RoleFindForUpdate(context.Background(), *req)
	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return &resp, nil
}
func (repo *AccountsRepo) RoleUpdate(ctx context.Context, req *db.RoleUpdateTXParams) (*db.AccountsSchemaRole, error) {
	resp, err := repo.store.RoleUpdateTX(context.Background(), *req)
	if err != nil {
		return nil, repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return &resp.Role, nil
}

func (repo *AccountsRepo) RoleDeleteRestore(ctx context.Context, req []int32) error {
	err := repo.store.RoleDeleteRestore(context.Background(), req)
	if err != nil {
		return repo.store.DbErrorParser(err, repo.errorHandler)
	}

	return nil
}
