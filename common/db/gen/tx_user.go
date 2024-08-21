package db

import (
	"context"

	"github.com/rs/zerolog/log"
)

func (store *SQLStore) AuthUserIDFindByEmail(ctx context.Context, email string) (string, error) {

	var response string
	err := store.execTX(ctx, func(q *Queries) error {
		err := q.db.QueryRow(ctx, "SELECT id FROM auth.users WHERE email = $1", email).Scan(&response)
		if err != nil {
			log.Debug().Interface("err", err).Msg("from AuthUserIDFindByEmail")
			return err
		}
		return err
	})

	return response, err
}
