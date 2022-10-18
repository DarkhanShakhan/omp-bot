package online

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/commands/education/online/commands"
	service "github.com/ozonmp/omp-bot/internal/service/education/online"
)

type OnlineCommander interface {
	Help(inputMsg *tgbotapi.Message)
	Get(inputMsg *tgbotapi.Message)
	List(inputMsg *tgbotapi.Message)
	Delete(inputMsg *tgbotapi.Message)

	New(inputMsg *tgbotapi.Message)  // return error not implemented
	Edit(inputMsg *tgbotapi.Message) // return error not implemented
}

func NewOnlineCommander(bot *tgbotapi.BotAPI, service service.OnlineService) OnlineCommander {
	return commands.NewCommander(bot, service)
}
