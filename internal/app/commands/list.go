package commands

import (
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strings"
)

func (c *Commander) List(inputMessage *tgbotapi.Message) {
	outputMsg := strings.Builder{}
	outputMsg.WriteString("Here all the products: \n\n")

	products := c.productService.List()
	for _, p := range products {
		outputMsg.WriteString(p.Title + "\n")
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsg.String())

	serializedData, _ := json.Marshal(CommandData{ //TODO HANDLE ERROR
		Offset: 21,
	})

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Next page", string(serializedData)),
		),
	)

	_, err := c.bot.Send(msg)

	if err != nil {
		log.Panic(err)
	}
}
