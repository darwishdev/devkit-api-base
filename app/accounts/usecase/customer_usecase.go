package usecase

import (
	"context"

	"github.com/bufbuild/connect-go"
	db "github.com/darwishdev/devkit-api-base/common/db/gen"
	abcv1 "github.com/darwishdev/devkit-api-base/common/pb/abc/v1"
	"github.com/darwishdev/devkit-api-base/common/supaapi"
	"github.com/rs/zerolog/log"
)

func (u *AccountsUsecase) CustomerLoginBase(ctx context.Context, req *abcv1.CustomerLoginRequest) (*abcv1.CustomerLoginResponse, error) {
	if err := u.validator.Validate(req); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	record, err := u.repo.CustomerFind(ctx, db.CustomerFindParams{
		SearchKey: req.LoginCode,
	})
	if err != nil {
		return nil, err
	}
	customerResp, err := u.adapter.CustomerLoginGrpcFromSql(record)
	if err != nil {
		return nil, err
	}
	return customerResp, nil
}
func (u *AccountsUsecase) CustomerLogin(ctx context.Context, req *abcv1.CustomerLoginRequest) (*abcv1.CustomerLoginResponse, error) {
	resp, err := u.CustomerLoginBase(ctx, req)
	if err != nil {
		return nil, err
	}
	_, err = u.supaClient.SignIn(ctx, supaapi.UserCredentials{
		Email:    resp.Customer.CustomerEmail,
		Password: req.CustomerPassword,
	})
	if err != nil {
		log.Debug().Interface("erorr is", err).Msg("what is happening")
		return nil, err
	}

	accessToken, accessPayload, err := u.tokenMaker.CreateToken(resp.Customer.CustomerEmail, resp.Customer.CustomerId, u.tokenDuration)
	if err != nil {
		return nil, err
	}

	loginInfo := &abcv1.LoginInfo{
		AccessToken:          accessToken,
		AccessTokenExpiresAt: accessPayload.ExpiredAt.Format("2006-01-02 15:04:05"),
	}
	resp.LoginInfo = loginInfo
	return resp, nil
}
func (u *AccountsUsecase) CustomerResetPassword(ctx context.Context, req *abcv1.CustomerResetPasswordRequest) error {
	if err := u.validator.Validate(req); err != nil {
		return connect.NewError(connect.CodeInvalidArgument, err)
	}
	params := u.adapter.CustomerResetPasswordSqlFromGrpc(req)
	err := u.repo.CustomerResetPassword(ctx, params)
	if err != nil {
		return err
	}
	return nil
}

func (u *AccountsUsecase) CustomersList(ctx context.Context, req *abcv1.Empty) (*abcv1.CustomersListResponse, error) {
	if err := u.validator.Validate(req); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	record, err := u.repo.CustomersList(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := u.adapter.CustomersListGrpcFromSql(record)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (u *AccountsUsecase) CustomerDeleteRestore(ctx context.Context, req *abcv1.DeleteRestoreRequest) (*abcv1.Empty, error) {
	if err := u.validator.Validate(req); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	err := u.repo.CustomerDeleteRestore(ctx, req.Records)
	if err != nil {
		return nil, err
	}
	return &abcv1.Empty{}, nil
}

func (u *AccountsUsecase) CustomerCreateUpdate(ctx context.Context, req *abcv1.CustomerCreateUpdateRequest) (*abcv1.CustomerEntity, error) {
	if err := u.validator.Validate(req); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	params := u.adapter.CustomerCreateUpdateSqlFromGrpc(req)
	record, err := u.repo.CustomerCreateUpdate(ctx, params)
	if err != nil {
		return nil, err
	}

	if req.CustomerId == 0 {
		_, err := u.supaClient.SignUp(ctx, supaapi.UserCredentials{
			Email:    record.CustomerEmail,
			Password: req.CustomerPassword,
		})
		if err != nil {
			return nil, err
		}
	} else {
		if req.CustomerPassword != "" {
			id, err := u.repo.AuthUserIDFindByEmail(ctx, req.CustomerEmail)
			if err != nil {
				return nil, err
			}
			supaParams := u.adapter.UserUpdateByIDSupaFromGrpc(req.CustomerEmail, req.CustomerPassword)
			respp, err := u.supaClient.UserUpdateById(ctx, *id, *supaParams)

			log.Debug().Interface("supaparams", *respp).Msg("hola")
			// log.Debug().Interface("supaparams", resss).Msg("hola")
			if err != nil {
				return nil, err
			}
		}
	}
	return u.adapter.CustomerEntityGrpcFromSql(record), nil

}

func (u *AccountsUsecase) CustomerRegister(ctx context.Context, req *abcv1.CustomerCreateUpdateRequest) (*abcv1.CustomerLoginResponse, error) {
	_, err := u.CustomerCreateUpdate(ctx, req)
	if err != nil {
		return nil, err
	}
	response, err := u.CustomerLogin(ctx, &abcv1.CustomerLoginRequest{
		LoginCode:        req.CustomerEmail,
		CustomerPassword: req.CustomerPassword,
	})
	if err != nil {
		return nil, err
	}
	return response, nil

}
