package gapi

import (
	"context"

	"github.com/bufbuild/connect-go"
	abcv1 "github.com/darwishdev/devkit-api-base/common/pb/abc/v1"
)

func (api *Api) OwnerLogin(ctx context.Context, req *connect.Request[abcv1.OwnerLoginRequest]) (*connect.Response[abcv1.OwnerLoginResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.accountsUsecase.OwnerLogin(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (api *Api) OwnersList(ctx context.Context, req *connect.Request[abcv1.Empty]) (*connect.Response[abcv1.OwnersListResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.accountsUsecase.OwnersList(ctx)
	if err != nil {
		return nil, err
	}
	opts, err := api.GetAccessableActionsForGroup(req.Header(), "owners")
	if err != nil {
		return nil, err
	}
	resp.Options = opts
	return connect.NewResponse(resp), nil
}

func (api *Api) OwnersInputList(ctx context.Context, req *connect.Request[abcv1.Empty]) (*connect.Response[abcv1.InputListResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.accountsUsecase.OwnersInputList(ctx)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}
func (api *Api) OwnerAuthorize(ctx context.Context, req *connect.Request[abcv1.Empty]) (*connect.Response[abcv1.OwnerLoginResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	authorizedOwner, _, err := api.authorizeUser(req.Header())
	if err != nil {
		return nil, err
	}
	resp, err := api.accountsUsecase.OwnerAuthorize(ctx, authorizedOwner)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (api *Api) OwnerResetPassword(ctx context.Context, req *connect.Request[abcv1.OwnerResetPasswordRequest]) (*connect.Response[abcv1.Empty], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	err := api.accountsUsecase.OwnerResetPassword(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(&abcv1.Empty{}), nil
}

func (api *Api) OwnerDeleteRestore(ctx context.Context, req *connect.Request[abcv1.DeleteRestoreRequest]) (*connect.Response[abcv1.Empty], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.accountsUsecase.OwnerDeleteRestore(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (api *Api) OwnerCreate(ctx context.Context, req *connect.Request[abcv1.OwnerCreateRequest]) (*connect.Response[abcv1.OwnerEntity], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.accountsUsecase.OwnerCreate(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}
func (api *Api) OwnerFind(ctx context.Context, req *connect.Request[abcv1.FindRequest]) (*connect.Response[abcv1.OwnersViewEntity], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.accountsUsecase.OwnerFind(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}
func (api *Api) OwnerFindForUpdate(ctx context.Context, req *connect.Request[abcv1.FindRequest]) (*connect.Response[abcv1.OwnerUpdateRequest], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.accountsUsecase.OwnerFindForUpdate(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}
func (api *Api) OwnerUpdate(ctx context.Context, req *connect.Request[abcv1.OwnerUpdateRequest]) (*connect.Response[abcv1.OwnerEntity], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.accountsUsecase.OwnerUpdate(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}
