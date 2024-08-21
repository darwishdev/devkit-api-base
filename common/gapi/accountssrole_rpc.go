package gapi

import (
	"context"
	"fmt"

	"github.com/bufbuild/connect-go"
	abcv1 "github.com/darwishdev/devkit-api-base/common/pb/abc/v1"
)

func (api *Api) RoleCreate(ctx context.Context, req *connect.Request[abcv1.RoleCreateRequest]) (*connect.Response[abcv1.RoleCreateResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.accountsUsecase.RoleCreate(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (api *Api) RoleFindForUpdate(ctx context.Context, req *connect.Request[abcv1.FindRequest]) (*connect.Response[abcv1.RoleUpdateRequest], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	resp, err := api.accountsUsecase.RoleFindForUpdate(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}
func (api *Api) RoleUpdate(ctx context.Context, req *connect.Request[abcv1.RoleUpdateRequest]) (*connect.Response[abcv1.RoleUpdateResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	if req.Msg.RoleId == 1 {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("admin group cannot be edited"))
	}
	resp, err := api.accountsUsecase.RoleUpdate(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (api *Api) PermissionsList(ctx context.Context, req *connect.Request[abcv1.PermissionsListRequest]) (*connect.Response[abcv1.PermissionsListResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.accountsUsecase.PermissionsList(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (api *Api) RolesList(ctx context.Context, req *connect.Request[abcv1.RolesListRequest]) (*connect.Response[abcv1.RolesListResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.accountsUsecase.RolesList(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	opts, err := api.GetAccessableActionsForGroup(req.Header(), "roles")
	if err != nil {
		return nil, err
	}
	resp.Options = opts
	return connect.NewResponse(resp), nil
}

func (api *Api) RolesInputList(ctx context.Context, req *connect.Request[abcv1.RolesInputListRequest]) (*connect.Response[abcv1.RolesInputListResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.accountsUsecase.RolesInputList(ctx, req.Msg)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(resp), nil
}

func searchInts32(arr []int32, searchNumber int32) bool {
	for _, elem := range arr {
		if elem == searchNumber {
			return true
		}
	}
	return false
}

func (api *Api) RoleDeleteRestore(ctx context.Context, req *connect.Request[abcv1.DeleteRestoreRequest]) (*connect.Response[abcv1.RoleDeleteRestoreResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	if searchInts32(req.Msg.Records, 1) {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("admin group cannot be deleted"))
	}
	resp, err := api.accountsUsecase.RoleDeleteRestore(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}
