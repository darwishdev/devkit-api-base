package usecase

import (
	"context"

	"github.com/bufbuild/connect-go"
	abcv1 "github.com/darwishdev/devkit-api-base/common/pb/abc/v1"
)

func (s *PublicUsecase) SettingsUpdate(ctx context.Context, req *abcv1.SettingsUpdateRequest) error {
	if err := s.validator.Validate(req); err != nil {
		return connect.NewError(connect.CodeInvalidArgument, err)
	}
	params := s.adapter.SettingsUpdateSqlFromGrpc(req)
	err := s.repo.SettingsUpdate(ctx, params)
	if err != nil {
		return err
	}
	return nil

}

func (u *PublicUsecase) SettingsFindForUpdate(ctx context.Context, req *abcv1.SettingsFindForUpdateRequest) (*abcv1.SettingsFindForUpdateResponse, error) {
	settings, err := u.repo.SettingsFindForUpdate(ctx)

	if err != nil {
		return nil, err
	}
	resp := u.adapter.SettingsFindForUpdateGrpcFromSql(settings)

	return resp, nil
}
func (s *PublicUsecase) IconsInputList(ctx context.Context) (*abcv1.IconsInputListResponse, error) {
	icons, err := s.repo.IconsInputList(ctx)
	if err != nil {
		return nil, err
	}
	res := s.adapter.IconsInputListGrpcFromSql(icons)
	return res, nil
}
