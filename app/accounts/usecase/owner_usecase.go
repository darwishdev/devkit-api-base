package usecase

import (
	"context"

	"github.com/bufbuild/connect-go"
	"github.com/darwishdev/devkit-api-base/common/auth"
	abcv1 "github.com/darwishdev/devkit-api-base/common/pb/abc/v1"
	"github.com/darwishdev/devkit-api-base/common/supaapi"
)

func (u *AccountsUsecase) OwnerLogin(ctx context.Context, req *abcv1.OwnerLoginRequest) (*abcv1.OwnerLoginResponse, error) {
	if err := u.validator.Validate(req); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	// params := u.adapter.OwnerLoginSqlFromGrpc(req)
	record, err := u.repo.OwnerFindByEmailOrCode(ctx, req.OwnerEmailOrCode)
	if err != nil {
		return nil, err
	}

	resp, err := u.adapter.OwnersViewEntityGrpcFromSql(record)
	if err != nil {
		return nil, err
	}

	accessToken, accessPayload, err := u.tokenMaker.CreateToken(resp.OwnerEmail, resp.OwnerId, u.tokenDuration)
	if err != nil {
		return nil, err
	}

	loginInfo := &abcv1.LoginInfo{
		AccessToken:          accessToken,
		AccessTokenExpiresAt: accessPayload.ExpiredAt.Format("2006-01-02 15:04:05"),
	}
	response := &abcv1.OwnerLoginResponse{
		LoginInfo: loginInfo,
		Owner:     resp,
	}
	return response, nil
}
func (u *AccountsUsecase) OwnerFind(ctx context.Context, req *abcv1.FindRequest) (*abcv1.OwnersViewEntity, error) {
	record, err := u.repo.OwnerFind(ctx, req.RecordId)
	if err != nil {
		return nil, err
	}
	response, err := u.adapter.OwnersViewEntityGrpcFromSql(record)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (u *AccountsUsecase) OwnerAuthorize(ctx context.Context, req *auth.Payload) (*abcv1.OwnerLoginResponse, error) {
	resp, err := u.OwnerLogin(ctx, &abcv1.OwnerLoginRequest{OwnerEmailOrCode: req.Username})
	return resp, err
}

func (u *AccountsUsecase) OwnerResetPassword(ctx context.Context, req *abcv1.OwnerResetPasswordRequest) error {
	params := u.adapter.OwnerResetPasswordSqlFromGrpc(req)
	err := u.repo.OwnerResetPassword(ctx, params)
	if err != nil {
		return err
	}
	return nil
}

func (u *AccountsUsecase) OwnersList(ctx context.Context) (*abcv1.OwnersListResponse, error) {
	record, err := u.repo.OwnersList(ctx)
	if err != nil {
		return nil, err
	}

	resp, err := u.adapter.OwnersListGrpcFromSql(*record)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (u *AccountsUsecase) OwnersInputList(ctx context.Context) (*abcv1.InputListResponse, error) {
	records, err := u.repo.OwnersInputList(ctx)
	if err != nil {
		return nil, err
	}
	resp := u.adapter.OwnersInputListGrpcFromSql(records)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (u *AccountsUsecase) OwnerDeleteRestore(ctx context.Context, req *abcv1.DeleteRestoreRequest) (*abcv1.Empty, error) {
	err := u.repo.OwnerDeleteRestore(ctx, req.Records)
	if err != nil {
		return nil, err
	}
	return &abcv1.Empty{}, nil
}

func (u *AccountsUsecase) OwnerCreate(ctx context.Context, req *abcv1.OwnerCreateRequest) (*abcv1.OwnerEntity, error) {
	if err := u.validator.Validate(req); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	params := u.adapter.OwnerCreateSqlFromGrpc(req)
	record, err := u.repo.OwnerCreate(ctx, params)
	if err != nil {
		return nil, err
	}

	_, err = u.supaClient.SignUp(ctx, supaapi.UserCredentials{
		Email:    record.OwnerEmail,
		Password: req.OwnerPassword,
	})
	if err != nil {
		return nil, err
	}

	return u.adapter.OwnerEntityGrpcFromSql(record), nil

}
func (u *AccountsUsecase) OwnerFindForUpdate(ctx context.Context, req *abcv1.FindRequest) (*abcv1.OwnerUpdateRequest, error) {
	user, err := u.repo.OwnerFindForUpdate(ctx, &req.RecordId)
	if err != nil {
		return nil, err
	}
	res := u.adapter.OwnerFindForUpdateGrpcFromSql(user)

	return res, nil
}

func (u *AccountsUsecase) OwnerUpdate(ctx context.Context, req *abcv1.OwnerUpdateRequest) (*abcv1.OwnerEntity, error) {
	if err := u.validator.Validate(req); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	params := u.adapter.OwnerUpdateSqlFromGrpc(req)
	record, err := u.repo.OwnerUpdate(ctx, params)
	if err != nil {
		return nil, err
	}
	if req.OwnerPassword != "" {
		id, err := u.repo.AuthUserIDFindByEmail(ctx, req.OwnerEmail)
		if err != nil {
			return nil, err
		}
		supaParams := u.adapter.UserUpdateByIDSupaFromGrpc(req.OwnerEmail, req.OwnerPassword)
		_, err = u.supaClient.UserUpdateById(ctx, *id, *supaParams)
		if err != nil {
			return nil, err
		}
	}
	// fmt.Println(supaOwner)
	return u.adapter.OwnerEntityGrpcFromSql(record), nil

}
