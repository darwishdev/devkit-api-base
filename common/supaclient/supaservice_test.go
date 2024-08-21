package supaclient

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/darwishdev/devkit-api-base/common/convertor"
	"github.com/darwishdev/devkit-api-base/common/random"
	"github.com/darwishdev/devkit-api-base/common/supaapi"
	// "github.com/darwishdev/devkit-api-base/common/"
)

type SigneUpTest struct {
	name      string
	req       *supaapi.UserCredentials
	expectErr bool
}

func getValidSignupRequest(field string, value interface{}) *supaapi.UserCredentials {
	validRole := &supaapi.UserCredentials{
		Email:    random.RandomEmail(),
		Password: random.RandomString(6),
	}
	if field != "" {
		err := convertor.SetField(validRole, field, value)
		if err != nil {
			log.Fatal(err)
		}
	}
	return validRole
}

func TestSignUp(t *testing.T) {
	// Define a slice of test cases
	testcases := []SigneUpTest{
		{
			name:      "ValidUser",
			req:       getValidSignupRequest("", ""),
			expectErr: false,
		},
	}
	fmt.Println("user  ")

	// Loop through the test cases and test each one
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {

			// Call the RoleCreate function with the role data from the current test case
			user, err := service.SignUp(context.Background(), *tc.req)

			fmt.Println("user  " + user.User.ID)
			// If the current test case expects an error and no error occurred, fail the test
			if tc.expectErr && err == nil {
				t.Errorf("Expected an error but got none")
			}

			// If the current test case does not expect an error and an error occurred, fail the test
			if !tc.expectErr && err != nil {
				t.Errorf("Expected no error but got %v", err)
			}
			//delete all  roles created during test
			// testQueries.RoleDelete(context.Background(), createdRole.RoleID)

		})
	}
}
