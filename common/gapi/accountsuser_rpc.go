package gapi

import (
	"context"

	"github.com/bufbuild/connect-go"
	abcv1 "github.com/darwishdev/devkit-api-base/common/pb/abc/v1"
)

func (api *Api) UserLogin(ctx context.Context, req *connect.Request[abcv1.UserLoginRequest]) (*connect.Response[abcv1.UserLoginResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.accountsUsecase.UserLogin(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (api *Api) UsersList(ctx context.Context, req *connect.Request[abcv1.Empty]) (*connect.Response[abcv1.UsersListResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.accountsUsecase.UsersList(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	opts, err := api.GetAccessableActionsForGroup(req.Header(), "users")
	if err != nil {
		return nil, err
	}
	resp.Options = opts
	return connect.NewResponse(resp), nil
}

func (api *Api) UserAuthorize(ctx context.Context, req *connect.Request[abcv1.Empty]) (*connect.Response[abcv1.Empty], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	_, _, err := api.authorizeUser(req.Header())
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&abcv1.Empty{}), nil
}

func (api *Api) UserResetPassword(ctx context.Context, req *connect.Request[abcv1.UserResetPasswordRequest]) (*connect.Response[abcv1.Empty], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	err := api.accountsUsecase.UserResetPassword(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(&abcv1.Empty{}), nil
}

func (api *Api) UserDeleteRestore(ctx context.Context, req *connect.Request[abcv1.DeleteRestoreRequest]) (*connect.Response[abcv1.Empty], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.accountsUsecase.UserDeleteRestore(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (api *Api) UserCreateUpdate(ctx context.Context, req *connect.Request[abcv1.UserCreateUpdateRequest]) (*connect.Response[abcv1.UserEntity], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.accountsUsecase.UserCreateUpdate(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (api *Api) UserFindForUpdate(ctx context.Context, req *connect.Request[abcv1.FindRequest]) (*connect.Response[abcv1.UserCreateUpdateRequest], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.accountsUsecase.UserFindForUpdate(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}
