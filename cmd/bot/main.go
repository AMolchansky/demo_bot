package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const helpCommand = "help"
const listCommand = "list"

func main() {
	godotenv.Load()

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
		if update.Message == nil {
			continue
		}

		switch update.Message.Command() {
		case helpCommand:
			processHelpCommand(bot, update.Message)
		case listCommand:
			processListCommand(bot, update.Message)
		default:
			processDefaultBehavior(bot, update.Message)
		}

		if update.Message.Command() == "help" {

			continue
		}

	}
}

func processHelpCommand(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
	sendMessage(bot, inputMessage.Chat.ID, ""+
		"/help - help\n"+
		"/list - list products",
	)
}

func processListCommand(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
	sendMessage(bot, inputMessage.Chat.ID, "TBD")
}

func processDefaultBehavior(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
	log.Printf("[%s] %s", inputMessage.From.UserName, inputMessage.Text)

	sendMessage(bot, inputMessage.Chat.ID, "You wrote: "+inputMessage.Text)
}

func sendMessage(bot *tgbotapi.BotAPI, chatID int64, text string) {
	msg := tgbotapi.NewMessage(chatID, text)

	_, err := bot.Send(msg)

	if err != nil {
		log.Panic(err)
	}
}
