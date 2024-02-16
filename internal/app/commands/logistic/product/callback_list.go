package product

import (
	"encoding/json"
	"fmt"
	"github.com/AMolchansky/demo_bot/internal/app/path"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

type CallbackListData struct {
	Offset int `json:"offset"`
}

func (pc *ProductCommander) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	parsedData := CallbackListData{}

	err := json.Unmarshal([]byte(callbackPath.CallbackData), &parsedData)
	if err != nil {
		log.Printf("ProductCommander.CallbackList: "+
			"error reading json data for type CallbackListData from "+
			"input string %v - %v", callbackPath.CallbackData, err)
		return
	}

	msg := tgbotapi.NewMessage(
		callback.Message.Chat.ID,
		fmt.Sprintf("Parsed: %+v\n", parsedData),
	)

	_, err = pc.bot.Send(msg)
	if err != nil {
		log.Printf("ProductCommander.CallbackList: error sending reply message to chat - %v", err)
	}
}
