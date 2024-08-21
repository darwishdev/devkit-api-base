package db

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/darwishdev/devkit-api-base/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

var store Store
var connPool *pgxpool.Pool

func intInSlice(target int32, slice []int32) bool {
	for _, item := range slice {
		if item == target {
			return true
		}
	}
	return false
}

func TestMain(m *testing.M) {
	ctx := context.Background()
	_, err := config.LoadState("../../../config")
	if err != nil {
		panic(err)
	}
	config, err := config.LoadConfig("../../../config", "test")

	if err != nil {
		fmt.Println("cannot load config")
		panic(err)
	}
	dbTracer := NewDbTracer()
	dbConfig, err := pgxpool.ParseConfig(config.DBSource)
	if err != nil {
		panic(err)
	}

	dbConfig.ConnConfig.Tracer = dbTracer
	connPool, err = pgxpool.NewWithConfig(ctx, dbConfig)
	if err != nil {
		panic(err)
	}
	// Attempt to ping the database to ensure the connection is working.
	if err := connPool.Ping(ctx); err != nil {
		connPool.Close()
		panic(err)
	}

	store = NewStore(connPool)

	os.Exit(m.Run())
}
