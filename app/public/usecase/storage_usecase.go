package usecase

import (
	"context"
	"fmt"

	"github.com/bufbuild/connect-go"
	"github.com/rs/zerolog/log"

	abcv1 "github.com/darwishdev/devkit-api-base/common/pb/abc/v1"
)

func (s *PublicUsecase) UploadFile(ctx context.Context, req *abcv1.UploadFileRequest) (*abcv1.UploadFileResponse, error) {
	if err := s.validator.Validate(req); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	params := s.adapter.UploadFileSupaClientFromGrpc(req)
	supaClientResponse, err := s.supaClient.Upload(*params)
	log.Debug().Interface("supResp", supaClientResponse).Interface("req", req).Msg("supRespsupRespsupRespsupResp")
	if err != nil {
		return nil, err
	}
	response := s.adapter.UploadFileGrpcFromSupaClient(supaClientResponse)
	if supaClientResponse.Message == "The resource already exists" {
		response.Path = fmt.Sprintf("images/%s", req.Path)
	}

	return response, nil

}

func (s *PublicUsecase) UploadFiles(ctx context.Context, req *abcv1.UploadFilesRequest) (*abcv1.UploadFileResponse, error) {
	if err := s.validator.Validate(req); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	params := s.adapter.UploadFilesSupaClientFromGrpc(req)
	supaClientResponse, err := s.supaClient.UploadMultiple(*params)
	if err != nil {
		return nil, err
	}

	response := s.adapter.UploadFilesGrpcFromSupaClient(supaClientResponse)

	return response, nil

}
