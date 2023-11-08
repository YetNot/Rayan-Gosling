package main

import (
	"Rayan-Gosling/internal/config"
	rayan "Rayan-Gosling/internal/rayan/images/db"
	"Rayan-Gosling/pkg/client"
	"context"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	cfg := config.GetConfig()

	postrgeSQLClient, err := client.NewClient(context.TODO(), 3, cfg.Storage)
	if err != nil {
		log.Fatalf("%v", err)
	}
	repository := rayan.NewRepository(postrgeSQLClient)

	if err := repository.InsertTable(context.TODO(), postrgeSQLClient); err != nil {
		log.Fatal(err)
	}

	TgBot()

}

func TgBot() {
	godotenv.Load()
	token := os.Getenv("TOKEN")
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.UpdateConfig{
		Timeout: 60,
	}

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		defer func() {
			if panicValue := recover(); panicValue != nil {
				log.Printf("recovered from panic: %v", panicValue)
			}
		}()

		if update.CallbackQuery != nil {
			msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Data: "+update.CallbackQuery.Data)
			bot.Send(msg)
			return
		}

	}
}