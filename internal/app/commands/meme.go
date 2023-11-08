package commands

import (
	//	rayan "Rayan-Gosling/internal/rayan/images/db"
	rayan "Rayan-Gosling/internal/rayan/images/db"
	"time"
	// "fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Meme(inputMessage *tgbotapi.Message, img *rayan.Image2) {
	for i := 0; i < 6; i++ {
		photoFileBytes := tgbotapi.FileBytes{
			Name:  "picture",
			Bytes: img.ImageData[i],
		}

		photo := tgbotapi.NewPhoto(inputMessage.Chat.ID, photoFileBytes)

		c.bot.Send(photo)

		time.Sleep(10 * time.Second)
	}
}
