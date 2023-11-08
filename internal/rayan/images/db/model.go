package rayan
import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Image1 struct {
	ImageID   []int
	ImageData [][]byte
}

type Image2 struct {
	ImageData [][]byte
}

type Repository interface {
	ReceivingTable(ctx context.Context, pool *pgxpool.Pool, img *Image2) error
	InsertTable(ctx context.Context, pool *pgxpool.Pool, img *Image1) error
}
