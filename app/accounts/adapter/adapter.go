package adapter

import (
	db "github.com/darwishdev/devkit-api-base/common/db/gen"
	abcv1 "github.com/darwishdev/devkit-api-base/common/pb/abc/v1"
	"github.com/darwishdev/devkit-api-base/common/supaapi"
)

type AccountsAdapterInterface interface {
	// entieis
	PermissionEntityGrpcFromSql(resp *db.AccountsSchemaPermission) *abcv1.PermissionEntity
	RoleEntityGrpcFromSql(resp *db.AccountsSchemaRole) *abcv1.RoleEntity
	RolePermissionEntityGrpcFromSql(resp *db.AccountsSchemaRolePermission) *abcv1.RolePermissionEntity
	UserEntityGrpcFromSql(resp *db.AccountsSchemaUser) *abcv1.UserEntity
	UserRoleEntityGrpcFromSql(resp *db.AccountsSchemaUserRole) *abcv1.UserRoleEntity
	UserPermissionEntityGrpcFromSql(resp *db.AccountsSchemaUserPermission) *abcv1.UserPermissionEntity
	NavigationBarEntityGrpcFromSql(resp *db.AccountsSchemaNavigationBar) *abcv1.NavigationBarEntity
	OwnerEntityGrpcFromSql(resp *db.AccountsSchemaOwner) *abcv1.OwnerEntity
	OwnersViewEntityGrpcFromSql(resp *db.AccountsSchemaOwnersView) (*abcv1.OwnersViewEntity, error)
	CustomerEntityGrpcFromSql(resp *db.AccountsSchemaCustomer) *abcv1.CustomerEntity
	//roles

	PermissionsListGrpcFromSql(resp []db.PermissionsListRow) (*abcv1.PermissionsListResponse, error)
	rolesListRowGrpcFromSql(resp *db.RolesListRow) *abcv1.RolesListRow
	RolesListGrpcFromSql(resp *[]db.RolesListRow) *abcv1.RolesListResponse
	RoleCreateSqlFromGrpc(req *abcv1.RoleCreateRequest) *db.RoleCreateTXParams
	RoleUpdateSqlFromGrpc(req *abcv1.RoleUpdateRequest) *db.RoleUpdateTXParams
	RoleUpdateGrpcFromSql(resp *db.AccountsSchemaRole) *abcv1.RoleUpdateResponse
	RoleFindForUpdateGrpcFromSql(resp *[]byte) (*abcv1.RoleUpdateRequest, error)
	RoleCreateGrpcFromSql(resp *db.AccountsSchemaRole) *abcv1.RoleCreateResponse
	RolesInputListGrpcFromSql(resp *[]db.RolesInputListRow) *abcv1.RolesInputListResponse

	//userUserFindByEmailOrCode(ctx context.Context, req string) (*db.UserFindByEmailOrCodeRow, error)
	UsersListGrpcFromSql(resp *[]db.UsersListRow) (*abcv1.UsersListResponse, error)
	UsersPermissionsMapFromSql(resp *[]db.UserPermissionsListRow) (map[string]map[string]bool, error)
	UserResetPasswordSqlFromGrpc(req *abcv1.UserResetPasswordRequest) *db.UserResetPasswordParams
	UserCreateUpdateSqlFromGrpc(req *abcv1.UserCreateUpdateRequest) *db.UserCreateUpdateParams
	UserFindForUpdateGrpcFromSql(resp *db.UserFindForUpdateRow) *abcv1.UserCreateUpdateRequest
	UserFindNavigationBarsGrpcFromSql(resp *[]db.UserFindNavigationBarsRow) (*[]*abcv1.SideBarItem, error)
	UserRolesListGrpcFromSql(resp *[]db.UserRolesListRow) (*[]*abcv1.UserRole, error)
	UserUpdateByIDSupaFromGrpc(email string, password string) *supaapi.AdminUserParams
	// UsersPermissionsMapFromSql(resp []db.UserPermissionsListRow) (*map[string]map[string]bool, error)
	//customers
	CustomersListGrpcFromSql(resp *[]db.AccountsSchemaCustomer) (*abcv1.CustomersListResponse, error)
	CustomerResetPasswordSqlFromGrpc(req *abcv1.CustomerResetPasswordRequest) *db.CustomerResetPasswordParams
	CustomerCreateUpdateSqlFromGrpc(req *abcv1.CustomerCreateUpdateRequest) *db.CustomerCreateUpdateParams
	CustomerLoginGrpcFromSql(resp *db.CustomerFindRow) (*abcv1.CustomerLoginResponse, error)
	// owners
	OwnersListGrpcFromSql(resp []db.AccountsSchemaOwner) (*abcv1.OwnersListResponse, error)
	OwnersInputListGrpcFromSql(resp *[]db.OwnersInputListRow) *abcv1.InputListResponse
	OwnerResetPasswordSqlFromGrpc(req *abcv1.OwnerResetPasswordRequest) *db.OwnerResetPasswordParams
	OwnerCreateSqlFromGrpc(req *abcv1.OwnerCreateRequest) *db.OwnerCreateParams
	OwnerUpdateSqlFromGrpc(req *abcv1.OwnerUpdateRequest) *db.OwnerUpdateParams
	OwnerFindForUpdateGrpcFromSql(resp *db.OwnerFindForUpdateRow) *abcv1.OwnerUpdateRequest
}

type AccountsAdapter struct {
	dateFormat string
}

func NewAccountsAdapter() AccountsAdapterInterface {
	return &AccountsAdapter{
		dateFormat: "2006-01-02 15:04:05",
	}
}
