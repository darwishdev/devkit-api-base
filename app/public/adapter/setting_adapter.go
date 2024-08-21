package adapter

import (
	db "github.com/darwishdev/devkit-api-base/common/db/gen"
	abcv1 "github.com/darwishdev/devkit-api-base/common/pb/abc/v1"
)

func (a *PublicAdapter) SettingsUpdateSqlFromGrpc(req *abcv1.SettingsUpdateRequest) *db.SettingsUpdateParams {
	keys := make([]string, len(req.Settings))
	values := make([]string, len(req.Settings))
	for index, v := range req.Settings {
		keys[index] = v.SettingKey
		values[index] = v.SettingValue
	}
	return &db.SettingsUpdateParams{
		Keys:   keys,
		Values: values,
	}
}
func (a *PublicAdapter) SettingsEntityGrpcFromSql(resp []db.Setting) []*abcv1.Setting {
	grpcResp := make([]*abcv1.Setting, len(resp))
	for _, v := range resp {
		record := &abcv1.Setting{
			SettingKey:   v.SettingKey,
			SettingValue: v.SettingValue,
		}
		grpcResp = append(grpcResp, record)
	}
	return grpcResp

}

func (a *PublicAdapter) SettingsFindForUpdateGrpcFromSql(resp *[]db.SettingsFindForUpdateRow) *abcv1.SettingsFindForUpdateResponse {
	grpcRows := make([]*abcv1.SettingsFindForUpdateRow, len(*resp))
	for index, v := range *resp {
		grpcRow := &abcv1.SettingsFindForUpdateRow{
			SettingKey:   v.SettingKey,
			SettingValue: v.SettingValue,
			SettingType:  v.SettingType,
		}

		grpcRows[index] = grpcRow

	}

	return &abcv1.SettingsFindForUpdateResponse{
		Settings: grpcRows,
	}

}

func (a *PublicAdapter) IconsInputListGrpcFromSql(resp *[]db.Icon) *abcv1.IconsInputListResponse {
	// IconsInputListGrpcFromSql
	records := make([]*abcv1.IconsInputListRow, 0)
	for _, v := range *resp {
		record := &abcv1.IconsInputListRow{
			IconId:      v.IconID,
			IconName:    v.IconName,
			IconContent: v.IconContent,
		}
		records = append(records, record)
	}
	return &abcv1.IconsInputListResponse{
		Icons: records,
	}
}
