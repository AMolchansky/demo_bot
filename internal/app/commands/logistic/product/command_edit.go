package product

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
	"strings"
)

func (pc *DummyProductCommander) Edit(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	argParts := strings.SplitN(args, " ", 2)
	if len(argParts) != 2 {
		pc.sendMessage(inputMessage.Chat.ID, "Please provide both product ID and title", "ProductCommander.Edit")
		return
	}

	productIdStr := argParts[0]
	newTitle := argParts[1]

	productId, err := strconv.Atoi(productIdStr)
	if err != nil {
		pc.sendMessage(inputMessage.Chat.ID, "Invalid product id", "ProductCommander.Edit")
		return
	}

	product, err := pc.dummyProductService.Describe(uint64(productId))
	if err != nil {
		pc.sendMessage(inputMessage.Chat.ID, "Product not found", "ProductCommander.Edit")
		return
	}

	if product.Title == newTitle {
		pc.sendMessage(inputMessage.Chat.ID, "You try to change title to the same", "ProductCommander.Edit")
		return
	}

	product.Title = newTitle

	err = pc.dummyProductService.Update(uint64(productId), *product)
	if err != nil {
		pc.sendMessage(inputMessage.Chat.ID, "Failed to update product", "ProductCommander.Edit")
		return
	}

	pc.sendMessage(inputMessage.Chat.ID, "Product updated successfully", "ProductCommander.Edit")
}
