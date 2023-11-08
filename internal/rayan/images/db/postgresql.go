package rayan

import (
	"Rayan-Gosling/pkg/client"
	"context"
	"os"
	"strconv"

	"github.com/jackc/pgx/v4/pgxpool"
)

type repository struct {
	client client.Client
}

func (r *repository) InsertTable(ctx context.Context, pool *pgxpool.Pool, img *Image1) (err error) {
	for i := 0; i < 6; i++ {
		a := strconv.Itoa(i+1)
		img.ImageData[i], err = os.ReadFile("C:/Users/Kirill/Desktop/rayan-gosling/" + a + ".jpg")
		if err != nil {
			return err
		}

		query := `
		INSERT INTO images (image) VALUES ($1) RETURNING id
	`
		err = pool.QueryRow(ctx, query, img.ImageData[i]).Scan(&img.ImageID[i])
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *repository) ReceivingTable(ctx context.Context, pool *pgxpool.Pool, img *Image2) error {
	for i := 1; i <= 6; i++ {
		query := "SELECT image FROM images WHERE id = $1"
		err := pool.QueryRow(ctx, query, i).Scan(&img.ImageData[i-1])
		if err != nil {
			return err
		}
	}

	return nil
}

func NewRepository(client client.Client) Repository {
	return &repository{
		client: client,
	}
}
