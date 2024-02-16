package product

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func (pc *DummyProductCommander) Default(inputMessage *tgbotapi.Message) {
	log.Printf("[%s] %s", inputMessage.From.UserName, inputMessage.Text)

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "You wrote: "+inputMessage.Text)

	_, err := pc.bot.Send(msg)

	if err != nil {
		log.Printf("ProductCommander.Help: error sending reply message to chat - %v", err)
	}
}
