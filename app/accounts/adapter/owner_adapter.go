package adapter

import (
	"github.com/darwishdev/devkit-api-base/common/convertor"
	db "github.com/darwishdev/devkit-api-base/common/db/gen"
	abcv1 "github.com/darwishdev/devkit-api-base/common/pb/abc/v1"

	"golang.org/x/crypto/bcrypt"
)

func (a *AccountsAdapter) OwnersListGrpcFromSql(resp []db.AccountsSchemaOwner) (*abcv1.OwnersListResponse, error) {
	records := make([]*abcv1.OwnerEntity, 0)
	deletedRecords := make([]*abcv1.OwnerEntity, 0)
	for _, v := range resp {
		record := a.OwnerEntityGrpcFromSql(&v)

		if record.DeletedAt == "" || record.DeletedAt == "0001-01-01 00:00:00" {
			records = append(records, record)
		} else {
			deletedRecords = append(deletedRecords, record)
		}

	}

	return &abcv1.OwnersListResponse{
		Records:        records,
		DeletedRecords: deletedRecords,
	}, nil
}

func (a *AccountsAdapter) OwnerResetPasswordSqlFromGrpc(req *abcv1.OwnerResetPasswordRequest) *db.OwnerResetPasswordParams {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.OwnerPassword), bcrypt.DefaultCost)

	return &db.OwnerResetPasswordParams{
		OwnerEmail:    req.OwnerEmail,
		OwnerPassword: string(hashedPassword),
	}

}

func (a *AccountsAdapter) OwnerCreateSqlFromGrpc(req *abcv1.OwnerCreateRequest) *db.OwnerCreateParams {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.OwnerPassword), bcrypt.DefaultCost)
	return &db.OwnerCreateParams{
		OwnerName:             req.OwnerName,
		OwnerImage:            convertor.ToPgType(req.OwnerImage),
		OwnerEmail:            req.OwnerEmail,
		RepresentativeOwnerID: convertor.ToPgTypeID(req.RepresentativeOwnerId),
		OwnerPhone:            convertor.ToPgType(req.OwnerPhone),
		OwnerPassword:         string(hashedPassword),
		OwnerNationalID:       req.OwnerNationalId,
	}
}

func (a *AccountsAdapter) OwnerUpdateSqlFromGrpc(req *abcv1.OwnerUpdateRequest) *db.OwnerUpdateParams {

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.OwnerPassword), bcrypt.DefaultCost)

	return &db.OwnerUpdateParams{
		OwnerID:               req.OwnerId,
		OwnerName:             req.OwnerName,
		OwnerImage:            convertor.ToPgType(req.OwnerImage),
		RepresentativeOwnerID: convertor.ToPgTypeID(req.RepresentativeOwnerId),
		OwnerEmail:            req.OwnerEmail,
		OwnerPhone:            convertor.ToPgType(req.OwnerPhone),
		OwnerPassword:         string(hashedPassword),
		OwnerNationalID:       req.OwnerNationalId,
	}
}

func (a *AccountsAdapter) OwnerFindForUpdateGrpcFromSql(resp *db.OwnerFindForUpdateRow) *abcv1.OwnerUpdateRequest {

	return &abcv1.OwnerUpdateRequest{
		OwnerId:               resp.OwnerID,
		OwnerName:             resp.OwnerName,
		OwnerImage:            resp.OwnerImage.String,
		RepresentativeOwnerId: resp.RepresentativeOwnerID.Int32,
		OwnerEmail:            resp.OwnerEmail,
		OwnerPhone:            resp.OwnerPhone.String,
		OwnerNationalId:       resp.OwnerNationalID}
}

func (a *AccountsAdapter) OwnersInputListGrpcFromSql(resp *[]db.OwnersInputListRow) *abcv1.InputListResponse {
	// OwnersInputListGrpcFromSql
	records := make([]*abcv1.SelectInputOption, 0)
	for _, v := range *resp {
		record := convertor.ToSelectInput(v.RecordValue, v.RecordLabel, "properties", v.Note)
		records = append(records, record)
	}
	return &abcv1.InputListResponse{
		Options: records,
	}
}
