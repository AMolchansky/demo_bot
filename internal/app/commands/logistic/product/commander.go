package product

import (
	"github.com/AMolchansky/demo_bot/internal/app/path"
	service "github.com/AMolchansky/demo_bot/internal/service/logistic/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

type ProductCommander interface {
	Help(inputMsg *tgbotapi.Message)
	Get(inputMsg *tgbotapi.Message)
	List(inputMsg *tgbotapi.Message)
	Delete(inputMsg *tgbotapi.Message)
	New(inputMsg *tgbotapi.Message)
	Edit(inputMsg *tgbotapi.Message)
}

type DummyProductCommander struct {
	bot                 *tgbotapi.BotAPI
	dummyProductService *service.DummyProductService
}

func NewProductCommander(
	bot *tgbotapi.BotAPI,
) *DummyProductCommander {
	dummyProductService := service.NewDummyProductService()

	return &DummyProductCommander{
		bot:                 bot,
		dummyProductService: dummyProductService,
	}
}

func (pc *DummyProductCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		pc.CallbackList(callback, callbackPath)
	default:
		log.Printf("ProductCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (pc *DummyProductCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		pc.Help(msg)
	case "list":
		pc.List(msg)
	case "get":
		pc.Get(msg)
	case "new":
		pc.New(msg)
	case "delete":
		pc.Delete(msg)
	case "edit":
		pc.Edit(msg)
	default:
		pc.Default(msg)
	}
}
