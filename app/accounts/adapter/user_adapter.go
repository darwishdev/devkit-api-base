package adapter

import (
	"encoding/json"

	db "github.com/darwishdev/devkit-api-base/common/db/gen"
	abcv1 "github.com/darwishdev/devkit-api-base/common/pb/abc/v1"
	"github.com/darwishdev/devkit-api-base/common/supaapi"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/rs/zerolog/log"

	"golang.org/x/crypto/bcrypt"
)

func (a *AccountsAdapter) UserUpdateByIDSupaFromGrpc(email string, password string) *supaapi.AdminUserParams {
	return &supaapi.AdminUserParams{
		Email:    email,
		Password: &password,
	}

}
func (a *AccountsAdapter) UsersListGrpcFromSql(resp *[]db.UsersListRow) (*abcv1.UsersListResponse, error) {
	records := make([]*abcv1.UserEntity, 0)
	deletedRecords := make([]*abcv1.UserEntity, 0)
	for _, v := range *resp {
		preventDelete := false
		for _, roleId := range v.RoleIds {
			log.Debug().Interface("new", roleId).Interface("user", v.UserID).Msg("hola")
			if roleId == 1 {
				preventDelete = true
				break
			}
		}
		record := &abcv1.UserEntity{
			UserId:        v.UserID,
			UserName:      v.UserName,
			UserImage:     v.UserImage.String,
			UserEmail:     v.UserEmail,
			UserPhone:     v.UserPhone.String,
			UserPassword:  v.UserPassword,
			PreventDelete: preventDelete,
			PreventUpdate: preventDelete,
			CreatedAt:     v.CreatedAt.Time.Format(a.dateFormat),
			UpdatedAt:     v.UpdatedAt.Time.Format(a.dateFormat),
			DeletedAt:     v.DeletedAt.Time.Format(a.dateFormat),
		}
		if v.DeletedAt.Valid {
			deletedRecords = append(deletedRecords, record)
		} else {
			records = append(records, record)
		}
	}
	return &abcv1.UsersListResponse{
		Records:        records,
		DeletedRecords: deletedRecords,
	}, nil
}

func (a *AccountsAdapter) UsersPermissionsMapFromSql(resp *[]db.UserPermissionsListRow) (map[string]map[string]bool, error) {
	response := make(map[string]map[string]bool, 0)
	for _, group := range *resp {
		if response[group.PermissionGroup] == nil {
			log.Debug().Interface("hoo", group).Msg("permissionsMap")
			response[group.PermissionGroup] = make(map[string]bool, 0)
		}
		response[group.PermissionGroup][group.PermissionFunction] = true

	}

	return response, nil
}

func (a *AccountsAdapter) UserResetPasswordSqlFromGrpc(req *abcv1.UserResetPasswordRequest) *db.UserResetPasswordParams {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.UserPassword), bcrypt.DefaultCost)

	return &db.UserResetPasswordParams{
		UserEmail:    req.UserEmail,
		UserPassword: string(hashedPassword),
	}

}

func (a *AccountsAdapter) UserCreateUpdateSqlFromGrpc(req *abcv1.UserCreateUpdateRequest) *db.UserCreateUpdateParams {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.UserPassword), bcrypt.DefaultCost)
	return &db.UserCreateUpdateParams{
		UserID:       req.UserId,
		UserName:     req.UserName,
		UserImage:    req.UserImage,
		UserEmail:    req.UserEmail,
		UserPhone:    req.UserPhone,
		UserPassword: string(hashedPassword),
		Permissions:  req.Permissions,
		Roles:        req.Roles,
	}
}

func (a *AccountsAdapter) UserFindForUpdateGrpcFromSql(resp *db.UserFindForUpdateRow) *abcv1.UserCreateUpdateRequest {
	// hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(resp.UserPassword), bcrypt.DefaultCost)
	return &abcv1.UserCreateUpdateRequest{
		UserId:      resp.UserID,
		UserName:    resp.UserName,
		UserImage:   resp.UserImage.String,
		UserEmail:   resp.UserEmail,
		UserCode:    resp.UserCode,
		UserPhone:   resp.UserPhone.String,
		Permissions: resp.Permissions,
		Roles:       resp.Roles,
	}
}
func (a *AccountsAdapter) UserFindNavigationBarsGrpcFromSql(resp *[]db.UserFindNavigationBarsRow) (*[]*abcv1.SideBarItem, error) {
	response := make([]*abcv1.SideBarItem, 0)
	for _, row := range *resp {
		items := make([]*abcv1.SideBarItem, len(row.Items))
		if len(row.Items) == 0 && !row.Route.Valid {
			continue
		}
		if len(row.Items) > 0 {
			if err := json.Unmarshal(row.Items, &items); err != nil {
				return nil, err
			}
		}

		if row.LabelAr.Valid {
			row.LabelAr = pgtype.Text{Valid: true, String: row.LabelAr.String}
		} else {
			row.LabelAr = pgtype.Text{Valid: true, String: row.Label}
		}
		response = append(response, &abcv1.SideBarItem{
			Key:     row.Key,
			Label:   row.Label,
			LabelAr: row.LabelAr.String,
			IconId:  row.IconID,
			Route:   row.Route.String,
			Items:   items,
		})
	}

	return &response, nil
}

func (a *AccountsAdapter) UserRolesListGrpcFromSql(resp *[]db.UserRolesListRow) (*[]*abcv1.UserRole, error) {
	response := make([]*abcv1.UserRole, len(*resp))
	for index, row := range *resp {
		response[index] = &abcv1.UserRole{
			RoleId:   row.RoleID,
			RoleName: row.RoleName,
		}
	}
	return &response, nil
}
