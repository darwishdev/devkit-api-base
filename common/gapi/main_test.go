package gapi

import (
	"testing"

	db "github.com/darwishdev/devkit-api-base/common/db/gen"
	"github.com/darwishdev/devkit-api-base/common/pb/abc/v1/abcv1connect"
	"github.com/darwishdev/devkit-api-base/common/redisclient"
	"github.com/darwishdev/devkit-api-base/config"
	"github.com/rs/zerolog/log"
)

func newTestApi(t *testing.T, store db.Store) abcv1connect.AbcServiceHandler {
	config, err := config.LoadConfig("./config", "test")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load the config")
	}
	rc := redisclient.NewRedisClient(config.RedisHost, config.RedisPort, config.RedisPassword, config.RedisDatabase)

	api := NewApi(config, store, rc)

	return api
}
