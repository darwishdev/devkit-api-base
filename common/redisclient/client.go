package redisclient

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type RedisClientInterface interface {
	AuthSessionCreate(ctx context.Context, userName string, permissions PermissionsMap) error
	AuthSessionFind(ctx context.Context, username string) (PermissionsMap, error)
}

type RedisClient struct {
	client *redis.Client
}

func NewRedisClient(host string, port string, password string, db int) RedisClientInterface {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: password, // no password set
		DB:       db,       // use default DB
	})

	return &RedisClient{
		client: client,
	}
}
