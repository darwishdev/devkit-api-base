package repo

import (
	"context"

	db "github.com/darwishdev/devkit-api-base/common/db/gen"
)

type AccountsRepoInterface interface {
	// role
	PermissionsList(ctx context.Context) ([]db.PermissionsListRow, error)
	RoleCreate(ctx context.Context, req *db.RoleCreateTXParams) (*db.AccountsSchemaRole, error)
	RoleUpdate(ctx context.Context, req *db.RoleUpdateTXParams) (*db.AccountsSchemaRole, error)
	RoleFindForUpdate(ctx context.Context, req *int32) (*[]byte, error)
	RolesList(ctx context.Context) (*[]db.RolesListRow, error)
	RoleDeleteRestore(ctx context.Context, req []int32) error
	RolesInputList(ctx context.Context) (*[]db.RolesInputListRow, error)
	// user
	UserFind(ctx context.Context, req db.UserFindParams) (*db.AccountsSchemaUser, error)
	UserPermissionsList(ctx context.Context, req int32) (*[]db.UserPermissionsListRow, error)
	UsersList(ctx context.Context) (*[]db.UsersListRow, error)
	UserResetPassword(ctx context.Context, req *db.UserResetPasswordParams) error
	UserDeleteRestore(ctx context.Context, req []int32) error
	UserCreateUpdate(ctx context.Context, req *db.UserCreateUpdateParams) (*db.AccountsSchemaUser, error)
	UserFindForUpdate(ctx context.Context, req *int32) (*db.UserFindForUpdateRow, error)
	UserFindNavigationBars(ctx context.Context, req *int32) (*[]db.UserFindNavigationBarsRow, error)
	UserRolesList(ctx context.Context, req *int32) (*[]db.UserRolesListRow, error)
	AuthUserIDFindByEmail(ctx context.Context, req string) (*string, error)
	// ownsers
	OwnerFindByEmailOrCode(ctx context.Context, req string) (*db.AccountsSchemaOwnersView, error)
	OwnersInputList(ctx context.Context) (*[]db.OwnersInputListRow, error)
	OwnerFind(ctx context.Context, req int32) (*db.AccountsSchemaOwnersView, error)
	OwnersList(ctx context.Context) (*[]db.AccountsSchemaOwner, error)
	OwnerResetPassword(ctx context.Context, req *db.OwnerResetPasswordParams) error
	OwnerDeleteRestore(ctx context.Context, req []int32) error
	OwnerCreate(ctx context.Context, req *db.OwnerCreateParams) (*db.AccountsSchemaOwner, error)
	OwnerFindForUpdate(ctx context.Context, req *int32) (*db.OwnerFindForUpdateRow, error)
	OwnerUpdate(ctx context.Context, req *db.OwnerUpdateParams) (*db.AccountsSchemaOwner, error)

	// customer

	CustomersList(ctx context.Context) (*[]db.AccountsSchemaCustomer, error)
	CustomerResetPassword(ctx context.Context, req *db.CustomerResetPasswordParams) error
	CustomerDeleteRestore(ctx context.Context, req []int32) error
	CustomerCreateUpdate(ctx context.Context, req *db.CustomerCreateUpdateParams) (*db.AccountsSchemaCustomer, error)
	CustomerFind(ctx context.Context, req db.CustomerFindParams) (*db.CustomerFindRow, error)
}

type AccountsRepo struct {
	store        db.Store
	errorHandler map[string]string
}

func NewAccountsRepo(store db.Store) AccountsRepoInterface {
	errorHandler := map[string]string{
		"roles_role_name_key":          "roleName",
		"users_user_name_key":          "userName",
		"owners_owner_phone_key":       "ownerPhone",
		"owners_owner_email_key":       "ownerEmail",
		"owners_owner_national_id_key": "ownerNationalId",
		"users_user_email_key":         "userEmail",
		"users_user_phone_key":         "userPhone",
		"customers_customer_email_key": "customerEmail",
		"customers_customer_phone_key": "customerPhone",
	}
	return &AccountsRepo{
		store:        store,
		errorHandler: errorHandler,
	}
}
