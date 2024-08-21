package redisclient

import (
	"context"
	"encoding/json"
)

type PermissionsMap map[string]map[string]bool

func (r *RedisClient) AuthSessionCreate(ctx context.Context, userName string, permissions PermissionsMap) error {
	jsonBytes, err := json.Marshal(permissions)
	if err != nil {
		return err
	}

	err = r.client.Set(ctx, userName, jsonBytes, 0).Err()
	if err != nil {
		return err
	}
	_, err = r.AuthSessionFind(ctx, userName)
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisClient) AuthSessionFind(ctx context.Context, username string) (PermissionsMap, error) {
	var parsedStruct PermissionsMap
	jsonBytes, err := r.client.Get(ctx, username).Bytes()
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(jsonBytes, &parsedStruct); err != nil {
		return nil, err
	}

	return parsedStruct, nil
}
