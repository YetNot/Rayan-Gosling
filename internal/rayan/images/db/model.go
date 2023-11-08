package rayan

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Image struct {
	imageID   int
	imageData []byte
}

type Repository interface {
	ReceivingTable(ctx context.Context, pool *pgxpool.Pool, img *Image) error
	InsertTable(ctx context.Context, pool *pgxpool.Pool) error
}
