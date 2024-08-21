package db

import (
	"context"
)

type RoleCreateTXParams struct {
	RoleParams        RoleCreateParams
	PermissionsParams []RolePermissionsBulkCreateParams
}

type RoleCreateTXResult struct {
	Role AccountsSchemaRole
}

func (store *SQLStore) RoleCreateTX(ctx context.Context, arg RoleCreateTXParams) (RoleCreateTXResult, error) {
	var result RoleCreateTXResult

	err := store.execTX(ctx, func(q *Queries) error {
		var err error

		role, err := q.RoleCreate(ctx, arg.RoleParams)
		if err != nil {
			return err
		}

		for i := 0; i < len(arg.PermissionsParams); i++ {
			arg.PermissionsParams[i].RoleID = role.RoleID
		}
		_, err = q.RolePermissionsBulkCreate(ctx, arg.PermissionsParams)
		if err != nil {
			return err
		}

		result.Role = role

		return err
	})

	return result, err
}

type RoleUpdateTXParams struct {
	RoleParams        RoleUpdateParams
	PermissionsParams []RolePermissionsBulkCreateParams
}

type RoleUpdateTXResult struct {
	Role AccountsSchemaRole
}

func (store *SQLStore) RoleUpdateTX(ctx context.Context, arg RoleUpdateTXParams) (RoleUpdateTXResult, error) {
	var result RoleUpdateTXResult

	err := store.execTX(ctx, func(q *Queries) error {
		var err error

		role, err := q.RoleUpdate(ctx, arg.RoleParams)
		if err != nil {
			return err
		}

		for i := 0; i < len(arg.PermissionsParams); i++ {
			arg.PermissionsParams[i].RoleID = role.RoleID
		}

		err = q.RolePermissionsClear(ctx, role.RoleID)
		if err != nil {
			return err
		}
		_, err = q.RolePermissionsBulkCreate(ctx, arg.PermissionsParams)
		if err != nil {
			return err
		}

		result.Role = role

		return err
	})

	return result, err
}
