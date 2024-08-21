package adapter

import (
	"encoding/json"

	db "github.com/darwishdev/devkit-api-base/common/db/gen"
	abcv1 "github.com/darwishdev/devkit-api-base/common/pb/abc/v1"
)

func (a *AccountsAdapter) PermissionEntityGrpcFromSql(resp *db.AccountsSchemaPermission) *abcv1.PermissionEntity {
	return &abcv1.PermissionEntity{
		PermissionId:          resp.PermissionID,
		PermissionFunction:    resp.PermissionFunction,
		PermissionName:        resp.PermissionName,
		PermissionDescription: resp.PermissionDescription.String,
		PermissionGroup:       resp.PermissionGroup,
	}
}
func (a *AccountsAdapter) RoleEntityGrpcFromSql(resp *db.AccountsSchemaRole) *abcv1.RoleEntity {
	return &abcv1.RoleEntity{
		RoleId:          resp.RoleID,
		RoleName:        resp.RoleName,
		RoleDescription: resp.RoleDescription.String,
		CreatedAt:       resp.CreatedAt.Time.Format(a.dateFormat),
		UpdatedAt:       resp.UpdatedAt.Time.Format(a.dateFormat),
		DeletedAt:       resp.DeletedAt.Time.Format(a.dateFormat),
	}
}
func (a *AccountsAdapter) RolePermissionEntityGrpcFromSql(resp *db.AccountsSchemaRolePermission) *abcv1.RolePermissionEntity {
	return &abcv1.RolePermissionEntity{
		RoleId:       resp.RoleID,
		PermissionId: resp.PermissionID,
	}
}

func (a *AccountsAdapter) UserEntityGrpcFromSql(resp *db.AccountsSchemaUser) *abcv1.UserEntity {
	return &abcv1.UserEntity{
		UserId:       resp.UserID,
		UserName:     resp.UserName,
		UserImage:    resp.UserImage.String,
		UserEmail:    resp.UserEmail,
		UserPhone:    resp.UserPhone.String,
		UserPassword: resp.UserPassword,
		CreatedAt:    resp.CreatedAt.Time.Format(a.dateFormat),
		UpdatedAt:    resp.UpdatedAt.Time.Format(a.dateFormat),
		DeletedAt:    resp.DeletedAt.Time.Format(a.dateFormat),
	}
}

func (a *AccountsAdapter) CustomerEntityGrpcFromSql(resp *db.AccountsSchemaCustomer) *abcv1.CustomerEntity {
	return &abcv1.CustomerEntity{
		CustomerId:         resp.CustomerID,
		CustomerCode:       resp.CustomerCode,
		CustomerName:       resp.CustomerName,
		CustomerImage:      resp.CustomerImage.String,
		Birthdate:          resp.Birthdate.Time.Format("2006-01-02"),
		CustomerEmail:      resp.CustomerEmail,
		CustomerPhone:      resp.CustomerPhone.String,
		CustomerPassword:   resp.CustomerPassword,
		CustomerNationalId: resp.CustomerNationalID.String,
		CreatedAt:          resp.CreatedAt.Time.Format(a.dateFormat),
		UpdatedAt:          resp.UpdatedAt.Time.Format(a.dateFormat),
		DeletedAt:          resp.DeletedAt.Time.Format(a.dateFormat),
	}
}
func (a *AccountsAdapter) UserRoleEntityGrpcFromSql(resp *db.AccountsSchemaUserRole) *abcv1.UserRoleEntity {
	return &abcv1.UserRoleEntity{
		UserId: resp.UserID,
		RoleId: resp.RoleID,
	}
}
func (a *AccountsAdapter) UserPermissionEntityGrpcFromSql(resp *db.AccountsSchemaUserPermission) *abcv1.UserPermissionEntity {
	return &abcv1.UserPermissionEntity{
		UserId:       resp.UserID,
		PermissionId: resp.PermissionID,
	}
}
func (a *AccountsAdapter) NavigationBarEntityGrpcFromSql(resp *db.AccountsSchemaNavigationBar) *abcv1.NavigationBarEntity {
	return &abcv1.NavigationBarEntity{
		NavigationBarId: resp.NavigationBarID,
		MenuKey:         resp.MenuKey,
		Label:           resp.Label,
		IconId:          resp.IconID,
		Route:           resp.Route.String,
		ParentId:        resp.ParentID.Int32,
		PermissionId:    resp.PermissionID.Int32,
	}
}

func (a *AccountsAdapter) OwnerEntityGrpcFromSql(resp *db.AccountsSchemaOwner) *abcv1.OwnerEntity {
	return &abcv1.OwnerEntity{
		OwnerId:         resp.OwnerID,
		OwnerName:       resp.OwnerName,
		OwnerImage:      resp.OwnerImage.String,
		OwnerEmail:      resp.OwnerEmail,
		OwnerPhone:      resp.OwnerPhone.String,
		OwnerNationalId: resp.OwnerNationalID,
		CreatedAt:       resp.CreatedAt.Time.Format(a.dateFormat),
		DeletedAt:       resp.DeletedAt.Time.Format(a.dateFormat),
	}
}
func (a *AccountsAdapter) OwnersViewEntityGrpcFromSql(resp *db.AccountsSchemaOwnersView) (*abcv1.OwnersViewEntity, error) {
	properties := make([]*abcv1.PropertiesViewEntity, 0)
	if len(resp.Properties) > 0 {
		if err := json.Unmarshal(resp.Properties, &properties); err != nil {
			return nil, err
		}
	}

	return &abcv1.OwnersViewEntity{
		OwnerId:         resp.OwnerID,
		OwnerName:       resp.OwnerName,
		OwnerImage:      resp.OwnerImage.String,
		OwnerEmail:      resp.OwnerEmail,
		OwnerPhone:      resp.OwnerPhone.String,
		OwnerNationalId: resp.OwnerNationalID,
		CreatedAt:       resp.CreatedAt.Time.Format(a.dateFormat),
		DeletedAt:       resp.DeletedAt.Time.Format(a.dateFormat),
		Properties:      properties,
	}, nil
}
