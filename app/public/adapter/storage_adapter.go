package adapter

import (
	"bytes"
	"io"
	"strings"

	abcv1 "github.com/darwishdev/devkit-api-base/common/pb/abc/v1"
	"github.com/darwishdev/devkit-api-base/common/supaapi"
	"github.com/darwishdev/devkit-api-base/common/supaclient"
)

func (a *PublicAdapter) UploadFileSupaClientFromGrpc(req *abcv1.UploadFileRequest) *supaclient.UploadRequest {

	reader := io.NopCloser(bytes.NewReader(req.Reader))
	return &supaclient.UploadRequest{
		BucketName: req.BucketName,
		Path:       req.Path,
		FileType:   req.FileType,
		Reader:     reader,
	}
}

func (a *PublicAdapter) UploadFilesSupaClientFromGrpc(req *abcv1.UploadFilesRequest) *supaclient.UploadMultipleRequest {

	records := make([]supaclient.UploadRequest, 0)
	for fileIndex, file := range req.Files {
		reader := io.NopCloser(bytes.NewReader(file.Reader))
		record := supaclient.UploadRequest{
			BucketName: file.BucketName,
			Path:       file.Path,
			FileType:   file.FileType,
			Reader:     reader,
		}

		records[fileIndex] = record

	}
	return &supaclient.UploadMultipleRequest{
		Files: records,
	}
}
func (a *PublicAdapter) UploadFileGrpcFromSupaClient(resp supaapi.FileResponse) *abcv1.UploadFileResponse {

	return &abcv1.UploadFileResponse{Path: resp.Key}
}
func (a *PublicAdapter) UploadFilesGrpcFromSupaClient(resp []*supaapi.FileResponse) *abcv1.UploadFileResponse {
	pathes := []string{}
	for _, file := range resp {
		pathes = append(pathes, file.Key)
	}
	path := strings.Join(pathes, ",")
	return &abcv1.UploadFileResponse{Path: path}
}
