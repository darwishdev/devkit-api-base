package adapter

import (
	"encoding/json"

	"github.com/darwishdev/devkit-api-base/common/convertor"
	db "github.com/darwishdev/devkit-api-base/common/db/gen"
	abcv1 "github.com/darwishdev/devkit-api-base/common/pb/abc/v1"
)

// func parsePermission(res []byte, resultch chan *abcv1.PermissionGroup, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	var record *abcv1.PermissionGroup
// 	err := json.Unmarshal(res, &record)
// 	if err != nil {
// 		// Handle the error, e.g., log it or send it to an error channel.
// 		return
// 	}

// 	resultch <- record
// }

func (a *AccountsAdapter) PermissionsListGrpcFromSql(resp []db.PermissionsListRow) (*abcv1.PermissionsListResponse, error) {
	response := make([]*abcv1.PermissionGroup, len(resp))
	// mu := sync.Mutex{}
	for index, v := range resp {
		var permissions []*abcv1.Permission
		err := json.Unmarshal(v.Permissions, &permissions)
		if err != nil {
			return nil, err
		}

		response[index] = &abcv1.PermissionGroup{
			PermissionGroup: v.PermissionGroup,
			Permissions:     permissions,
		}
	}
	return &abcv1.PermissionsListResponse{
		Records: response,
	}, nil
}

//list

func (a *AccountsAdapter) rolesListRowGrpcFromSql(resp *db.RolesListRow) *abcv1.RolesListRow {

	return &abcv1.RolesListRow{
		RoleId:           resp.RoleID,
		RoleName:         resp.RoleName,
		RoleDescription:  resp.RoleDescription.String,
		PreventDelete:    resp.RoleID == 1,
		PreventUpdate:    resp.RoleID == 1,
		PermissionsCount: int32(resp.PermissionsCount),
		UsersCount:       int32(resp.UsersCount),
		CreatedAt:        resp.CreatedAt.Time.Format(a.dateFormat),
		DeletedAt:        resp.DeletedAt.Time.Format(a.dateFormat),
	}
}

func (a *AccountsAdapter) RolesListGrpcFromSql(resp *[]db.RolesListRow) *abcv1.RolesListResponse {
	records := make([]*abcv1.RolesListRow, 0)
	deletedRecords := make([]*abcv1.RolesListRow, 0)
	for _, v := range *resp {
		record := a.rolesListRowGrpcFromSql(&v)
		if v.DeletedAt.Valid {
			deletedRecords = append(deletedRecords, record)
		} else {
			records = append(records, record)
		}
	}
	return &abcv1.RolesListResponse{
		Records:        records,
		DeletedRecords: deletedRecords,
	}
}

func (a *AccountsAdapter) RolesInputListGrpcFromSql(resp *[]db.RolesInputListRow) *abcv1.RolesInputListResponse {
	// RolesInputListGrpcFromSql
	records := make([]*abcv1.SelectInputOption, 0)
	for _, v := range *resp {
		record := convertor.ToSelectInput(v.RoleID, v.RoleName, "", "")
		records = append(records, record)
	}
	return &abcv1.RolesInputListResponse{
		Options: records,
	}
}

func (a *AccountsAdapter) RoleCreateSqlFromGrpc(req *abcv1.RoleCreateRequest) *db.RoleCreateTXParams {
	permissionsParams := make([]db.RolePermissionsBulkCreateParams, 0)
	for _, v := range req.Permissions {
		rolePermission := db.RolePermissionsBulkCreateParams{
			PermissionID: v,
		}
		permissionsParams = append(permissionsParams, rolePermission)
	}
	return &db.RoleCreateTXParams{
		RoleParams: db.RoleCreateParams{
			RoleName:        req.RoleName,
			RoleDescription: convertor.ToPgType(req.RoleDescription),
		},
		PermissionsParams: permissionsParams,
	}
}
func RoleEntityGrpcFromSql(resp *db.AccountsSchemaRole) *abcv1.Role {
	return &abcv1.Role{
		RoleId:          int32(resp.RoleID),
		RoleName:        resp.RoleName,
		RoleDescription: resp.RoleDescription.String,
		CreatedAt:       resp.CreatedAt.Time.Format("2006-01-02 15:04:05"),
		DeletedAt:       resp.DeletedAt.Time.Format("2006-01-02 15:04:05"),
	}

}
func (a *AccountsAdapter) RoleCreateGrpcFromSql(resp *db.AccountsSchemaRole) *abcv1.RoleCreateResponse {
	return &abcv1.RoleCreateResponse{
		Role: RoleEntityGrpcFromSql(resp),
	}
}

func (a *AccountsAdapter) RoleFindForUpdateGrpcFromSql(resp *[]byte) (*abcv1.RoleUpdateRequest, error) {
	var response abcv1.RoleUpdateRequest
	err := json.Unmarshal([]byte(*resp), &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (a *AccountsAdapter) RoleUpdateSqlFromGrpc(req *abcv1.RoleUpdateRequest) *db.RoleUpdateTXParams {
	permissionsParams := make([]db.RolePermissionsBulkCreateParams, 0)
	for _, v := range req.Permissions {
		rolePermission := db.RolePermissionsBulkCreateParams{
			PermissionID: v,
		}
		permissionsParams = append(permissionsParams, rolePermission)
	}
	return &db.RoleUpdateTXParams{
		RoleParams: db.RoleUpdateParams{
			RoleID:          req.RoleId,
			RoleName:        req.RoleName,
			RoleDescription: convertor.ToPgType(req.RoleDescription),
		},
		PermissionsParams: permissionsParams,
	}
}
func (a *AccountsAdapter) RoleUpdateGrpcFromSql(resp *db.AccountsSchemaRole) *abcv1.RoleUpdateResponse {
	return &abcv1.RoleUpdateResponse{
		Role: RoleEntityGrpcFromSql(resp),
	}
}
