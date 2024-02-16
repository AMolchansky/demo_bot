package product

import (
	"encoding/json"
	"github.com/AMolchansky/demo_bot/internal/app/path"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strings"
)

func (pc *ProductCommander) List(inputMessage *tgbotapi.Message) {
	outputMsg := strings.Builder{}
	outputMsg.WriteString("Here all the products: \n\n")

	products := pc.productService.List()
	for _, p := range products {
		outputMsg.WriteString(p.Title + "\n")
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsg.String())

	serializedData, _ := json.Marshal(CommandData{ //TODO HANDLE ERROR
		Offset: 21,
	})

	callbackPath := path.CallbackPath{
		Domain:       "logistic",
		Subdomain:    "product",
		CallbackName: "list",
		CallbackData: string(serializedData),
	}

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Next page", callbackPath.String()),
		),
	)

	_, err := pc.bot.Send(msg)

	if err != nil {
		log.Printf("ProductCommander.List: error sending reply message to chat - %v", err)
	}
}
