package main

import (
	"github.com/AMolchansky/demo_bot/internal/app/commands"
	"github.com/AMolchansky/demo_bot/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"log"
	"os"
)

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

	productService := product.NewService()

	commander := commands.NewCommander(bot, productService)

	for update := range updates {
		commander.HandleUpdate(update)
	}
}
