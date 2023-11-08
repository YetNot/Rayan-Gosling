package rayan

import (
	// rayan "Rayan-Gosling/internal/rayan/images/db"
	// "Rayan-Gosling/pkg/client"
	"Rayan-Gosling/pkg/client"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

type repository struct {
	client client.Client
}

func (r *repository) InsertTable(ctx context.Context, pool *pgxpool.Pool) (err error) {
	var img = &Image{}
	img.imageData, err = os.ReadFile("C:/Users/Kirill/Desktop/rayan-gosling/1.jpg")
	if err != nil {
		return err
	}

	query := `
	INSERT INTO images (image) VALUES ($1) RETURNING id
`
	err = pool.QueryRow(ctx, query, img.imageData).Scan(&img.imageID)
	if err != nil {
		return err
	}

	log.Printf("Image inserted with ID: %d\n", img.imageID)

	r.ReceivingTable(ctx, pool, img)

	return nil
}

func (r *repository) ReceivingTable(ctx context.Context, pool *pgxpool.Pool, img *Image) error {
	fmt.Println(img.imageID)
	query := "SELECT image FROM images WHERE id = $1"
	err := pool.QueryRow(ctx, query, img.imageID).Scan(&img.imageData)
	if err != nil {
		return err
	}

	err = os.WriteFile("C:/Users/Kirill/Desktop/1a.jpg", img.imageData, 0644)
	if err != nil {
		return err
	}
	return nil
}

func NewRepository(client client.Client) Repository {
	return &repository{
		client: client,
	}
}
