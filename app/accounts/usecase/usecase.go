package usecase

import (
	"context"
	"time"

	"github.com/bufbuild/protovalidate-go"
	"github.com/darwishdev/devkit-api-base/app/accounts/adapter"
	"github.com/darwishdev/devkit-api-base/app/accounts/repo"
	"github.com/darwishdev/devkit-api-base/common/auth"
	db "github.com/darwishdev/devkit-api-base/common/db/gen"
	abcv1 "github.com/darwishdev/devkit-api-base/common/pb/abc/v1"
	"github.com/darwishdev/devkit-api-base/common/redisclient"
	"github.com/darwishdev/devkit-api-base/common/resend"
	"github.com/darwishdev/devkit-api-base/common/supaclient"
)

type AccountsUsecaseInterface interface {
	RoleCreate(ctx context.Context, req *abcv1.RoleCreateRequest) (*abcv1.RoleCreateResponse, error)
	RoleFindForUpdate(ctx context.Context, req *abcv1.FindRequest) (*abcv1.RoleUpdateRequest, error)
	RoleUpdate(ctx context.Context, req *abcv1.RoleUpdateRequest) (*abcv1.RoleUpdateResponse, error)
	PermissionsList(ctx context.Context, req *abcv1.PermissionsListRequest) (*abcv1.PermissionsListResponse, error)
	RolesList(ctx context.Context, req *abcv1.RolesListRequest) (*abcv1.RolesListResponse, error)
	RoleDeleteRestore(ctx context.Context, req *abcv1.DeleteRestoreRequest) (*abcv1.RoleDeleteRestoreResponse, error)
	RolesInputList(ctx context.Context, req *abcv1.RolesInputListRequest) (*abcv1.RolesInputListResponse, error)

	// user
	UserLogin(ctx context.Context, req *abcv1.UserLoginRequest) (*abcv1.UserLoginResponse, error)
	UserResetPassword(ctx context.Context, req *abcv1.UserResetPasswordRequest) error
	UsersList(ctx context.Context, req *abcv1.Empty) (*abcv1.UsersListResponse, error)
	UserDeleteRestore(ctx context.Context, req *abcv1.DeleteRestoreRequest) (*abcv1.Empty, error)
	UserCreateUpdate(ctx context.Context, req *abcv1.UserCreateUpdateRequest) (*abcv1.UserEntity, error)
	UserFindForUpdate(ctx context.Context, req *abcv1.FindRequest) (*abcv1.UserCreateUpdateRequest, error)

	//customer s
	CustomerLogin(ctx context.Context, req *abcv1.CustomerLoginRequest) (*abcv1.CustomerLoginResponse, error)
	CustomerResetPassword(ctx context.Context, req *abcv1.CustomerResetPasswordRequest) error
	CustomersList(ctx context.Context, req *abcv1.Empty) (*abcv1.CustomersListResponse, error)
	CustomerDeleteRestore(ctx context.Context, req *abcv1.DeleteRestoreRequest) (*abcv1.Empty, error)
	CustomerCreateUpdate(ctx context.Context, req *abcv1.CustomerCreateUpdateRequest) (*abcv1.CustomerEntity, error)
	CustomerRegister(ctx context.Context, req *abcv1.CustomerCreateUpdateRequest) (*abcv1.CustomerLoginResponse, error)
	CustomerLoginBase(ctx context.Context, req *abcv1.CustomerLoginRequest) (*abcv1.CustomerLoginResponse, error)
	// owners
	OwnerLogin(ctx context.Context, req *abcv1.OwnerLoginRequest) (*abcv1.OwnerLoginResponse, error)
	OwnerAuthorize(ctx context.Context, req *auth.Payload) (*abcv1.OwnerLoginResponse, error)
	OwnerResetPassword(ctx context.Context, req *abcv1.OwnerResetPasswordRequest) error
	OwnersList(ctx context.Context) (*abcv1.OwnersListResponse, error)
	OwnerDeleteRestore(ctx context.Context, req *abcv1.DeleteRestoreRequest) (*abcv1.Empty, error)
	OwnerCreate(ctx context.Context, req *abcv1.OwnerCreateRequest) (*abcv1.OwnerEntity, error)
	OwnerFindForUpdate(ctx context.Context, req *abcv1.FindRequest) (*abcv1.OwnerUpdateRequest, error)
	OwnerUpdate(ctx context.Context, req *abcv1.OwnerUpdateRequest) (*abcv1.OwnerEntity, error)
	OwnersInputList(ctx context.Context) (*abcv1.InputListResponse, error)
	OwnerFind(ctx context.Context, req *abcv1.FindRequest) (*abcv1.OwnersViewEntity, error)
}

type AccountsUsecase struct {
	repo          repo.AccountsRepoInterface
	validator     *protovalidate.Validator
	tokenMaker    auth.Maker
	tokenDuration time.Duration
	adapter       adapter.AccountsAdapterInterface
	redisClient   redisclient.RedisClientInterface
	resendClient  resend.ResendServiceInterface
	supaClient    supaclient.SupabaseServiceInterface
	supaToken     string
}

func NewAccountsUsecase(store db.Store, validator *protovalidate.Validator, tokenMaker auth.Maker, tokenDuration time.Duration, redisClient redisclient.RedisClientInterface, supaClient supaclient.SupabaseServiceInterface, supaToken string, resendClient resend.ResendServiceInterface) AccountsUsecaseInterface {
	repo := repo.NewAccountsRepo(store)
	adapter := adapter.NewAccountsAdapter()
	return &AccountsUsecase{
		repo:          repo,
		tokenMaker:    tokenMaker,
		validator:     validator,
		tokenDuration: tokenDuration,
		resendClient:  resendClient,
		adapter:       adapter,
		redisClient:   redisClient,
		supaClient:    supaClient,
		supaToken:     supaToken,
	}
}
