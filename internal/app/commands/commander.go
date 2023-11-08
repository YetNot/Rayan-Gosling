package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type DummyCommander interface {
	Meme(inputMsg *tgbotapi.Message)
}

type Commander struct {
	bot *tgbotapi.BotAPI
}

func NewCommander(bot *tgbotapi.BotAPI) *Commander {
	return &Commander{
		bot: bot,
	}
}
