package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func (c *Commander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, ""+
		"/help - help\n"+
		"/list - list products\n"+
		"/get - get product info",
	)

	_, err := c.bot.Send(msg)

	if err != nil {
		log.Panic(err)
	}
}
