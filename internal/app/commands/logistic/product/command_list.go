package product

import (
	"encoding/json"
	"github.com/AMolchansky/demo_bot/internal/app/path"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strings"
)

func (pc *DummyProductCommander) List(inputMessage *tgbotapi.Message) {
	outputMsg := strings.Builder{}

	var cursor uint64 = 1
	var offset uint64 = 5

	products, errIndex := pc.dummyProductService.List(cursor, offset)

	if errIndex != nil {
		pc.sendMessage(inputMessage.Chat.ID, "Invalid page", "ProductCommander.Get")
		return
	}

	if len(products) == 0 {
		pc.sendMessage(inputMessage.Chat.ID, "Products not found", "ProductCommander.Get")
		return
	}

	for _, p := range products {
		outputMsg.WriteString(p.String())
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsg.String())

	// check is next page available
	_, errIndex = pc.dummyProductService.List(cursor+1, offset)
	if errIndex == nil {
		nextPageSerializedData, errMarshal := json.Marshal(CallbackListData{
			Page:   int(cursor + 1),
			Offset: int(offset),
		})

		if errMarshal != nil {
			log.Printf("ProductCommander.List: error json.Marshal next page CallbackListData - %v", errMarshal)
			return
		}

		nextPageCallbackPath := path.CallbackPath{
			Domain:       "logistic",
			Subdomain:    "product",
			CallbackName: "list",
			CallbackData: string(nextPageSerializedData),
		}

		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Next page", nextPageCallbackPath.String()),
			),
		)
	}

	_, errSend := pc.bot.Send(msg)

	if errSend != nil {
		log.Printf("ProductCommander.List: error sending reply message to chat - %v", errSend)
	}
}
