package logistic

import (
	"github.com/AMolchansky/demo_bot/internal/app/commands/logistic/product"
	"github.com/AMolchansky/demo_bot/internal/app/path"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

type ICommander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type LogisticCommander struct {
	bot              *tgbotapi.BotAPI
	productCommander ICommander
}

func NewLogisticCommander(
	bot *tgbotapi.BotAPI,
) *LogisticCommander {
	return &LogisticCommander{
		bot:              bot,
		productCommander: product.NewProductCommander(bot),
	}
}

func (c *LogisticCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "product":
		c.productCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("LogisticCommander.HandleCallback: unknown subdomain - %s", callbackPath.Subdomain)
	}
}

func (c *LogisticCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "product":
		c.productCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("LogisticCommander.HandleCommand: unknown subdomain - %s", commandPath.Subdomain)
	}
}
