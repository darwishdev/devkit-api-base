package adapter

import (
	db "github.com/darwishdev/devkit-api-base/common/db/gen"
	abcv1 "github.com/darwishdev/devkit-api-base/common/pb/abc/v1"
	"github.com/darwishdev/devkit-api-base/common/supaapi"
	"github.com/darwishdev/devkit-api-base/common/supaclient"
)

type PublicAdapterInterface interface {
	SettingsUpdateSqlFromGrpc(req *abcv1.SettingsUpdateRequest) *db.SettingsUpdateParams
	SettingsEntityGrpcFromSql(resp []db.Setting) []*abcv1.Setting
	SettingsFindForUpdateGrpcFromSql(resp *[]db.SettingsFindForUpdateRow) *abcv1.SettingsFindForUpdateResponse
	IconsInputListGrpcFromSql(resp *[]db.Icon) *abcv1.IconsInputListResponse

	UploadFileSupaClientFromGrpc(req *abcv1.UploadFileRequest) *supaclient.UploadRequest
	UploadFileGrpcFromSupaClient(resp supaapi.FileResponse) *abcv1.UploadFileResponse
	UploadFilesSupaClientFromGrpc(req *abcv1.UploadFilesRequest) *supaclient.UploadMultipleRequest
	UploadFilesGrpcFromSupaClient(resp []*supaapi.FileResponse) *abcv1.UploadFileResponse
}

type PublicAdapter struct {
	dateFormat string
}

func NewPublicAdapter() PublicAdapterInterface {
	return &PublicAdapter{
		dateFormat: "2006-01-02 15:04:05",
	}
}
