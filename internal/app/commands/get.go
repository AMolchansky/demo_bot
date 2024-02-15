package commands

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
)

func (c *Commander) Get(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		log.Println("Wrong arg", args)
		return
	}

	product, err := c.productService.Get(idx)
	if err != nil {
		log.Printf("Fail to get product with ids %d: %v", idx, err)
		return
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("Succesfully parsed argument: %v", product.Title))

	_, err = c.bot.Send(msg)

	if err != nil {
		log.Panic(err)
	}
}
