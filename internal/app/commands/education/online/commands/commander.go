package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	service "github.com/ozonmp/omp-bot/internal/service/education/online"
)

type Commander struct {
	bot     *tgbotapi.BotAPI
	service service.OnlineService
}

func NewCommander(bot *tgbotapi.BotAPI, service service.OnlineService) *Commander {
	return &Commander{bot: bot, service: service}
}
