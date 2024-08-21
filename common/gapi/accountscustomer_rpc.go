package gapi

import (
	"context"

	"github.com/bufbuild/connect-go"
	abcv1 "github.com/darwishdev/devkit-api-base/common/pb/abc/v1"
)

func (api *Api) CustomerLogin(ctx context.Context, req *connect.Request[abcv1.CustomerLoginRequest]) (*connect.Response[abcv1.CustomerLoginResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.accountsUsecase.CustomerLogin(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (api *Api) CustomersList(ctx context.Context, req *connect.Request[abcv1.Empty]) (*connect.Response[abcv1.CustomersListResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.accountsUsecase.CustomersList(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	opts, err := api.GetAccessableActionsForGroup(req.Header(), "customers")
	if err != nil {
		return nil, err
	}
	resp.Options = opts
	return connect.NewResponse(resp), nil
}

func (api *Api) CustomerAuthorize(ctx context.Context, req *connect.Request[abcv1.Empty]) (*connect.Response[abcv1.CustomerLoginResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	payload, err := api.authorizeCustomer(req.Header())
	if err != nil {
		return nil, err
	}
	resp, err := api.accountsUsecase.CustomerLoginBase(ctx, &abcv1.CustomerLoginRequest{
		LoginCode: payload.Username,
	})
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (api *Api) CustomerResetPassword(ctx context.Context, req *connect.Request[abcv1.CustomerResetPasswordRequest]) (*connect.Response[abcv1.Empty], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	err := api.accountsUsecase.CustomerResetPassword(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(&abcv1.Empty{}), nil
}

func (api *Api) CustomerDeleteRestore(ctx context.Context, req *connect.Request[abcv1.DeleteRestoreRequest]) (*connect.Response[abcv1.Empty], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.accountsUsecase.CustomerDeleteRestore(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (api *Api) CustomerCreateUpdate(ctx context.Context, req *connect.Request[abcv1.CustomerCreateUpdateRequest]) (*connect.Response[abcv1.CustomerEntity], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.accountsUsecase.CustomerCreateUpdate(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (api *Api) CustomerRegister(ctx context.Context, req *connect.Request[abcv1.CustomerCreateUpdateRequest]) (*connect.Response[abcv1.CustomerLoginResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.accountsUsecase.CustomerRegister(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}
