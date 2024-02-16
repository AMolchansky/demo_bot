package product

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func (pc *DummyProductCommander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, ""+
		"/help__logistic__product - help\n"+
		"/list__logistic__product - list products\n"+
		"/get__logistic__product - get product info",
	)

	_, err := pc.bot.Send(msg)

	if err != nil {
		log.Printf("ProductCommander.Help: error sending reply message to chat - %v", err)
	}
}
