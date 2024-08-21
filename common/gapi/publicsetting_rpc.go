package gapi

import (
	"context"

	"github.com/bufbuild/connect-go"
	abcv1 "github.com/darwishdev/devkit-api-base/common/pb/abc/v1"
)

func (api *Api) SettingsUpdate(ctx context.Context, req *connect.Request[abcv1.SettingsUpdateRequest]) (*connect.Response[abcv1.SettingsUpdateResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	err := api.publicUsecase.SettingsUpdate(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(&abcv1.SettingsUpdateResponse{}), nil
}

func (api *Api) SettingsFindForUpdate(ctx context.Context, req *connect.Request[abcv1.SettingsFindForUpdateRequest]) (*connect.Response[abcv1.SettingsFindForUpdateResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	resp, err := api.publicUsecase.SettingsFindForUpdate(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}
