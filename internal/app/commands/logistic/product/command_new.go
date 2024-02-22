package product

import (
	"fmt"
	"github.com/AMolchansky/demo_bot/internal/model/logistic"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strings"
)

func (pc *DummyProductCommander) New(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	fields := strings.Fields(args)
	if len(fields) == 0 {
		pc.sendMessage(
			inputMessage.Chat.ID,
			"No fields provided. Please specify at least one field for the product.",
			"ProductCommander.New",
		)
		return
	}

	product := logistic.Product{}
	for _, field := range fields {
		parts := strings.SplitN(field, ":", 2)
		if len(parts) == 2 {
			fieldName := strings.TrimSpace(parts[0])
			fieldValue := strings.TrimSpace(parts[1])
			switch fieldName {
			case "title":
				product.Title = fieldValue
			}
		} else {
			pc.sendMessage(
				inputMessage.Chat.ID,
				"Fields must be provided by parameter:value structure",
				"ProductCommander.New",
			)
			return
		}
	}

	productId, err := pc.dummyProductService.Create(product)

	if err != nil {
		log.Printf("ProductCommander.New: error creating new product - %v", err)
		return
	}

	pc.sendMessage(
		inputMessage.Chat.ID,
		fmt.Sprintf("Sucessfully created new product with id: %d", productId),
		"ProductCommander.New",
	)
}
