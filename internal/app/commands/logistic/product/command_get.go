package product

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
)

func (pc *DummyProductCommander) Get(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		log.Println("wrong args", args)
		return
	}

	product, err := pc.dummyProductService.Get(idx)
	if err != nil {
		log.Printf("fail to get product with ids %d: %v", idx, err)
		return
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, product.Title)

	_, err = pc.bot.Send(msg)
	if err != nil {
		log.Printf("ProductCommander.Get: error sending reply message to chat - %v", err)
	}
}
