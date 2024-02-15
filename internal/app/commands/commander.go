package commands

import (
	"encoding/json"
	"fmt"
	"github.com/AMolchansky/demo_bot/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

const helpCommand = "help"
const listCommand = "list"
const getCommand = "get"

type Commander struct {
	bot            *tgbotapi.BotAPI
	productService *product.Service
}

type CommandData struct {
	Offset int `json:"offset"`
}

func NewCommander(
	bot *tgbotapi.BotAPI,
	productService *product.Service,
) *Commander {
	return &Commander{
		bot:            bot,
		productService: productService,
	}
}

func (c *Commander) HandleUpdate(update tgbotapi.Update) {
	defer func() {
		if panicValue := recover(); panicValue != nil {
			log.Printf("recoved from panic: %v", panicValue)
		}
	}()

	if update.CallbackQuery != nil {
		parsedData := CommandData{}
		json.Unmarshal([]byte(update.CallbackQuery.Data), &parsedData)
		msg := tgbotapi.NewMessage(
			update.CallbackQuery.Message.Chat.ID,
			fmt.Sprintf("Parsed: %+v\n", parsedData),
		)

		_, err := c.bot.Send(msg)

		if err != nil {
			log.Panic(err)
		}

		return
	}

	if update.Message == nil {
		return
	}

	switch update.Message.Command() {
	case helpCommand:
		c.Help(update.Message)
	case listCommand:
		c.List(update.Message)
	case getCommand:
		c.Get(update.Message)
	default:
		c.Default(update.Message)
	}
}
