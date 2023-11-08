package storage

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

func CreateTable(ctx context.Context, pool *pgxpool.Pool) error {
	query := `
	CREATE TABLE IF NOT EXISTS images (
		id SERIAL PRIMARY KEY,
		image BYTEA
	)
`
	_, err := pool.Exec(ctx, query)
	if err != nil {
		return err
	}

	return nil
}
