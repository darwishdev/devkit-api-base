package gapi

import (
	"context"

	"github.com/bufbuild/connect-go"
	abcv1 "github.com/darwishdev/devkit-api-base/common/pb/abc/v1"
)

func (api *Api) UploadFile(ctx context.Context, req *connect.Request[abcv1.UploadFileRequest]) (*connect.Response[abcv1.UploadFileResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	resp, err := api.publicUsecase.UploadFile(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}
