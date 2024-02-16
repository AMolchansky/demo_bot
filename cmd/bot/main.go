package main

import (
	routerPkg "github.com/AMolchansky/demo_bot/internal/app/router"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	_ = godotenv.Load()

	token, found := os.LookupEnv("TOKEN")
	if !found {
		log.Panic("environment variable TOKEN not found in .env")
	}

	bot, errInit := tgbotapi.NewBotAPI(token)
	if errInit != nil {
		log.Panic(errInit)
	}

	// Uncomment if you want debugging
	// bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.UpdateConfig{
		Timeout: 60,
	}

	updates := bot.GetUpdatesChan(u)

	routerHandler := routerPkg.NewRouter(bot)

	for update := range updates {
		routerHandler.HandleUpdate(update)
	}
}
