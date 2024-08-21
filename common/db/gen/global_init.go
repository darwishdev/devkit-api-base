package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func InitDB(ctx context.Context, source string) (Store, error) {
	dbTracer := NewDbTracer()
	dbConfig, err := pgxpool.ParseConfig(source)
	if err != nil {
		return nil, err
	}

	dbConfig.ConnConfig.Tracer = dbTracer
	connPool, err := pgxpool.NewWithConfig(ctx, dbConfig)
	if err != nil {
		return nil, err
	}
	// Attempt to ping the database to ensure the connection is working.
	if err := connPool.Ping(ctx); err != nil {
		connPool.Close()
		return nil, err
	}
	fmt.Println("connected to", source)
	store := NewStore(connPool)

	return store, nil
}
