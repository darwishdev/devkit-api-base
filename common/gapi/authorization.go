// Copyright 2022 Buf Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gapi

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/darwishdev/devkit-api-base/common/auth"
	abcv1 "github.com/darwishdev/devkit-api-base/common/pb/abc/v1"
	"github.com/darwishdev/devkit-api-base/common/redisclient"
	"github.com/iancoleman/strcase"
	"github.com/tangzero/inflector"
)

const (
	authorizationHeader = "authorization"
	authorizationBearer = "bearer"
)

func (api *Api) authorizeUser(header http.Header) (*auth.Payload, redisclient.PermissionsMap, error) {
	authHeader := header.Get("Authorization")
	if authHeader == "" {
		return nil, nil, fmt.Errorf("missing metadata")
	}

	fields := strings.Fields(authHeader)
	if len(fields) < 2 {
		return nil, nil, fmt.Errorf("invalid authorization header format")
	}

	authType := strings.ToLower(fields[0])
	if authType != authorizationBearer {
		return nil, nil, fmt.Errorf("unsupported authorization type: %s", authType)
	}

	accessToken := fields[1]
	payload, err := api.tokenMaker.VerifyToken(accessToken)
	if err != nil {
		return nil, nil, fmt.Errorf("invalid access token: %s", err)
	}

	authSession, err := api.redisClient.AuthSessionFind(context.Background(), payload.Username)
	if err != nil {
		_, err := api.accountsUsecase.UserLogin(context.Background(), &abcv1.UserLoginRequest{LoginCode: payload.Username})
		if err != nil {
			return nil, nil, fmt.Errorf("canot get the cache: %s", err)
		}

		authSession, err = api.redisClient.AuthSessionFind(context.Background(), payload.Username)
		if err != nil {
			return nil, nil, fmt.Errorf("canot get the cache: %s", err)
		}
	}

	return payload, authSession, nil
}

func (api *Api) authorizeCustomer(header http.Header) (*auth.Payload, error) {
	authHeader := header.Get("Authorization")
	if authHeader == "" {
		return nil, fmt.Errorf("missing metadata")
	}

	fields := strings.Fields(authHeader)
	if len(fields) < 2 {
		return nil, fmt.Errorf("invalid authorization header format")
	}

	authType := strings.ToLower(fields[0])
	if authType != authorizationBearer {
		return nil, fmt.Errorf("unsupported authorization type: %s", authType)
	}

	accessToken := fields[1]
	payload, err := api.tokenMaker.VerifyToken(accessToken)
	if err != nil {
		return nil, fmt.Errorf("invalid access token: %s", err)
	}

	return payload, nil
}

func (api *Api) GetAccessableActionsForGroup(header http.Header, group string) (*abcv1.ListDataOptions, error) {
	resp := abcv1.ListDataOptions{
		Title:       fmt.Sprintf("%s_list", group),
		Description: fmt.Sprintf("%s_description", group),
	}
	var (
		singularizedGroup string = inflector.Singularize(group)
		redirectRoute     string = fmt.Sprintf("%s_list", group)
		// requestProperty              string = fmt.Sprintf("%sId", singularizedGroup)
		deleteRestoreRequestProperty string = "records"
		findEndpoint                 string = strcase.ToLowerCamel(fmt.Sprintf("%s_find_for_update", singularizedGroup))
		findRequestProperty          string = "recordId"
		create                       string = fmt.Sprintf("%s_create", singularizedGroup)
		update                       string = fmt.Sprintf("%s_update", singularizedGroup)
		deleteRestore                string = fmt.Sprintf("%s_delete_restore", singularizedGroup)
	)
	_, permissionsMap, err := api.authorizeUser(header)
	if err != nil {
		return nil, err
	}
	if permissionsMap == nil {
		return &resp, nil
	}

	authorities := permissionsMap[group]
	if len(authorities) == 0 {
		return &resp, nil
	}
	if authorities[strcase.ToCamel(create)] {
		resp.CreateHandler = &abcv1.CreateHandler{
			RedirectRoute: redirectRoute,
			Title:         create,
			Endpoint:      strcase.ToLowerCamel(create),
			RouteName:     create,
		}
		// resp.ImportHandler = &abcv1.ImportHandler{
		// 	Endpoint:           importEndpoint,
		// 	ImportTemplateLink: importTemplateLink,
		// }
	}
	if authorities[strcase.ToCamel(update)] {
		resp.UpdateHandler = &abcv1.UpdateHandler{
			RedirectRoute:       redirectRoute,
			Title:               update,
			Endpoint:            strcase.ToLowerCamel(update),
			RouteName:           update,
			FindEndpoint:        findEndpoint,
			FindRequestProperty: findRequestProperty,
		}
	}
	if authorities[strcase.ToCamel(deleteRestore)] {
		resp.DeleteRestoreHandler = &abcv1.DeleteRestoreHandler{
			Endpoint:        strcase.ToLowerCamel(deleteRestore),
			RequestProperty: deleteRestoreRequestProperty,
		}
	}
	return &resp, nil
}
