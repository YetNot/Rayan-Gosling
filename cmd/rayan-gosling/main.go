package main

import (
	"Rayan-Gosling/internal/config"
	rayan "Rayan-Gosling/internal/rayan/images/db"
	"Rayan-Gosling/pkg/client"
	"context"
	"log"
)

func main() {
	cfg := config.GetConfig()

	postrgeSQLClient, err := client.NewClient(context.TODO(), 3, cfg.Storage)
	if err != nil {
		log.Fatalf("%v", err)
	}
	repository := rayan.NewRepository(postrgeSQLClient)

	var img = &rayan.Image1{
		ImageID: make([]int, 6),
		ImageData: make([][]byte, 6),
	}

	if err := repository.InsertTable(context.TODO(), postrgeSQLClient, img); err != nil {
		log.Fatal(err)
	}
}
