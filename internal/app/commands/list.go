package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
)

func (c *Commander) List(inputMessage *tgbotapi.Message) {
	outputMsg := strings.Builder{}
	outputMsg.WriteString("Here all the products: \n\n")

	products := c.productService.List()
	for _, p := range products {
		outputMsg.WriteString(p.Title + "\n")
	}

	c.sendMessage(inputMessage.Chat.ID, outputMsg.String())
}
