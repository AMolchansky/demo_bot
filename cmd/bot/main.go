package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	token := os.Getenv("TOKEN")

	bot, errInit := tgbotapi.NewBotAPI(token)
	if errInit != nil {
		log.Panic(errInit)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.UpdateConfig{
		Timeout: 60,
	}

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "You wrote:"+update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID

			_, errSend := bot.Send(msg)
			if errSend != nil {
				log.Panic(errSend)
			}
		}
	}
}
