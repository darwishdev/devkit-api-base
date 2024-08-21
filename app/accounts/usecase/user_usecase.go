package usecase

import (
	"context"
	"fmt"

	"github.com/bufbuild/connect-go"
	db "github.com/darwishdev/devkit-api-base/common/db/gen"
	abcv1 "github.com/darwishdev/devkit-api-base/common/pb/abc/v1"
	"github.com/darwishdev/devkit-api-base/common/supaapi"
)

func (u *AccountsUsecase) UserLogin(ctx context.Context, req *abcv1.UserLoginRequest) (*abcv1.UserLoginResponse, error) {
	if err := u.validator.Validate(req); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	record, err := u.repo.UserFind(ctx, db.UserFindParams{
		SearchKey: req.LoginCode,
	})
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("loginCode"))
		// return nil, err
	}

	_, err = u.supaClient.SignIn(ctx, supaapi.UserCredentials{
		Email:    record.UserEmail,
		Password: req.UserPassword,
	})
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("userPassword"))
	}

	userResp := u.adapter.UserEntityGrpcFromSql(record)

	permissions, err := u.repo.UserPermissionsList(ctx, userResp.UserId)
	if err != nil {
		return nil, err
	}
	permissionsMap, err := u.adapter.UsersPermissionsMapFromSql(permissions)
	if err != nil {
		return nil, err
	}

	userNavigations, err := u.repo.UserFindNavigationBars(ctx, &userResp.UserId)
	if err != nil {
		return nil, err
	}
	sidebar, err := u.adapter.UserFindNavigationBarsGrpcFromSql(userNavigations)
	if err != nil {
		return nil, err
	}

	userRoles, err := u.repo.UserRolesList(ctx, &userResp.UserId)
	if err != nil {
		return nil, err
	}

	roles, err := u.adapter.UserRolesListGrpcFromSql(userRoles)
	if err != nil {
		return nil, err
	}
	accessToken, accessPayload, err := u.tokenMaker.CreateToken(userResp.UserEmail, userResp.UserId, u.tokenDuration)
	if err != nil {
		return nil, err
	}
	err = u.redisClient.AuthSessionCreate(ctx, userResp.UserEmail, permissionsMap)
	if err != nil {
		return nil, err
	}
	loginInfo := &abcv1.LoginInfo{
		AccessToken:          accessToken,
		AccessTokenExpiresAt: accessPayload.ExpiredAt.Format("2006-01-02 15:04:05"),
	}
	response := &abcv1.UserLoginResponse{
		LoginInfo: loginInfo,
		User:      userResp,
		SideBar:   *sidebar,
		Roles:     *roles,
	}
	return response, nil
}
func (u *AccountsUsecase) UserResetPassword(ctx context.Context, req *abcv1.UserResetPasswordRequest) error {
	params := u.adapter.UserResetPasswordSqlFromGrpc(req)
	err := u.repo.UserResetPassword(ctx, params)
	if err != nil {
		return err
	}
	return nil
}

func (s *AccountsUsecase) UsersList(ctx context.Context, req *abcv1.Empty) (*abcv1.UsersListResponse, error) {
	record, err := s.repo.UsersList(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := s.adapter.UsersListGrpcFromSql(record)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *AccountsUsecase) UserDeleteRestore(ctx context.Context, req *abcv1.DeleteRestoreRequest) (*abcv1.Empty, error) {
	err := s.repo.UserDeleteRestore(ctx, req.Records)
	if err != nil {
		return nil, err
	}
	return &abcv1.Empty{}, nil
}

func (u *AccountsUsecase) UserCreateUpdate(ctx context.Context, req *abcv1.UserCreateUpdateRequest) (*abcv1.UserEntity, error) {
	params := u.adapter.UserCreateUpdateSqlFromGrpc(req)
	record, err := u.repo.UserCreateUpdate(ctx, params)
	if err != nil {
		return nil, err
	}

	if req.UserId == 0 {
		_, err := u.supaClient.SignUp(ctx, supaapi.UserCredentials{
			Email:    record.UserEmail,
			Password: req.UserPassword,
		})
		if err != nil {
			return nil, err
		}
	} else {
		if req.UserPassword != "" {
			id, err := u.repo.AuthUserIDFindByEmail(ctx, req.UserEmail)
			if err != nil {
				return nil, err
			}
			supaParams := u.adapter.UserUpdateByIDSupaFromGrpc(req.UserEmail, req.UserPassword)
			_, err = u.supaClient.UserUpdateById(ctx, *id, *supaParams)
			if err != nil {
				return nil, err
			}
		}
	}
	return u.adapter.UserEntityGrpcFromSql(record), nil

}
func (u *AccountsUsecase) UserFindForUpdate(ctx context.Context, req *abcv1.FindRequest) (*abcv1.UserCreateUpdateRequest, error) {
	user, err := u.repo.UserFindForUpdate(ctx, &req.RecordId)

	if err != nil {
		return nil, err
	}
	res := u.adapter.UserFindForUpdateGrpcFromSql(user)

	return res, nil
}
