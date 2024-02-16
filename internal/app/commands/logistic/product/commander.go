package product

import (
	"github.com/AMolchansky/demo_bot/internal/app/path"
	"github.com/AMolchansky/demo_bot/internal/service/logistic/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

type ProductCommander struct {
	bot            *tgbotapi.BotAPI
	productService *product.Service
}

type CommandData struct {
	Offset int `json:"offset"`
}

func NewProductCommander(
	bot *tgbotapi.BotAPI,
) *ProductCommander {
	productService := product.NewService()

	return &ProductCommander{
		bot:            bot,
		productService: productService,
	}
}

func (pc *ProductCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		pc.CallbackList(callback, callbackPath)
	default:
		log.Printf("ProductCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (pc *ProductCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		pc.Help(msg)
	case "list":
		pc.List(msg)
	case "get":
		pc.Get(msg)
	default:
		pc.Default(msg)
	}
}
