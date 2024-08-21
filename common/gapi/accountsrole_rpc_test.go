package gapi

import (
	"context"
	"testing"

	"github.com/bufbuild/connect-go"
	"github.com/darwishdev/devkit-api-base/common/convertor"
	db "github.com/darwishdev/devkit-api-base/common/db/gen"
	mockdb "github.com/darwishdev/devkit-api-base/common/db/mock"
	abcv1 "github.com/darwishdev/devkit-api-base/common/pb/abc/v1"
	"github.com/darwishdev/devkit-api-base/common/random"
	"github.com/golang/mock/gomock"
)

type roleCreateTest struct {
	name       string
	params     *abcv1.RoleCreateRequest
	buildStubs func(store *mockdb.MockStore)
	expectErr  bool
}

func getValidRole() *abcv1.RoleCreateRequest {
	return &abcv1.RoleCreateRequest{
		RoleName:        random.RandomString(10),
		RoleDescription: random.RandomString(50),
		Permissions:     []int32{random.RandomInt32(1, 60), random.RandomInt32(1, 60)},
	}
}

func TestRoleCreate(t *testing.T) {

	validRole := getValidRole()
	// Define a slice of test cases
	testcases := []roleCreateTest{
		// Test for a valid role creation.
		{
			name: "ValidRole",
			params: &abcv1.RoleCreateRequest{
				RoleName:        validRole.RoleName,
				RoleDescription: validRole.RoleDescription,
				Permissions:     validRole.Permissions,
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.RoleCreateTXParams{
					RoleParams: db.RoleCreateParams{
						RoleName:        validRole.RoleName,
						RoleDescription: convertor.ToPgType(validRole.RoleDescription),
					},
					PermissionsParams: []db.RolePermissionsBulkCreateParams{
						{
							PermissionID: validRole.Permissions[0],
						},
						{
							PermissionID: validRole.Permissions[1],
						},
					},
				}
				store.EXPECT().
					RoleCreateTX(gomock.Any(), arg).
					Times(1).
					Return(db.RoleCreateTXResult{Role: db.AccountsSchemaRole{
						RoleID:          1,
						RoleName:        validRole.RoleName,
						RoleDescription: convertor.ToPgType(validRole.RoleDescription),
					}}, nil)

			},
			expectErr: false,
		},
		{
			name: "InValidNameToShort",
			params: &abcv1.RoleCreateRequest{
				RoleName:        random.RandomString(1),
				RoleDescription: validRole.RoleDescription,
				Permissions:     validRole.Permissions,
			},
			buildStubs: func(store *mockdb.MockStore) {

				store.EXPECT().
					RoleCreateTX(gomock.Any(), gomock.Any()).
					Times(0)

			},
			expectErr: true,
		},
		{
			name: "InValidNameToLong",
			params: &abcv1.RoleCreateRequest{
				RoleName:        random.RandomString(220),
				RoleDescription: validRole.RoleDescription,
				Permissions:     validRole.Permissions,
			},
			buildStubs: func(store *mockdb.MockStore) {

				store.EXPECT().
					RoleCreateTX(gomock.Any(), gomock.Any()).
					Times(0)

			},
			expectErr: true,
		},
		{
			name: "InValidDescriptionToLong",
			params: &abcv1.RoleCreateRequest{
				RoleName:        random.RandomString(120),
				RoleDescription: random.RandomString(220),
				Permissions:     validRole.Permissions,
			},
			buildStubs: func(store *mockdb.MockStore) {

				store.EXPECT().
					RoleCreateTX(gomock.Any(), gomock.Any()).
					Times(0)

			},
			expectErr: true,
		},
		{
			name: "InvalideDuplicatedPermissions",
			params: &abcv1.RoleCreateRequest{
				RoleName:        random.RandomString(120),
				RoleDescription: random.RandomString(22),
				Permissions:     []int32{1, 1},
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					RoleCreateTX(gomock.Any(), gomock.Any()).
					Times(0)

			},
			expectErr: true,
		},
	}

	// Loop through the test cases and test each one
	// ctx := context.Background()
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			storeCtrl := gomock.NewController(t)
			defer storeCtrl.Finish()
			store := mockdb.NewMockStore(storeCtrl)
			tc.buildStubs(store)
			api := newTestApi(t, store)
			createdRole, err := api.RoleCreate(context.Background(), connect.NewRequest[abcv1.RoleCreateRequest](tc.params))
			// If the current test case expects an error and no error occurred, fail the test
			if tc.expectErr && err == nil {
				t.Errorf("Expected an error but got none %s", tc.name)
			}

			// If the current test case does not expect an error and an error occurred, fail the test
			if !tc.expectErr && err != nil {
				t.Errorf("Expected no error but got %v", err)
			}
			if !tc.expectErr {
				if createdRole.Msg.Role.RoleName != tc.params.RoleName {
					t.Errorf("un expected name wanted %s got %s", createdRole.Msg.Role.RoleName, tc.params.RoleName)
				}
				if createdRole.Msg.Role.RoleDescription != tc.params.RoleDescription {
					t.Errorf("un expected description wanted %s got %s", createdRole.Msg.Role.RoleDescription, tc.params.RoleDescription)
				}

			}

		})
	}
}
