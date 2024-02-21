package product

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
)

func (pc *DummyProductCommander) Edit(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	productId, err := strconv.Atoi(args)
	if err != nil {
		pc.sendMessage(inputMessage.Chat.ID, "Invalid product id", "ProductCommander.Edit")
		return
	}

	product, err := pc.dummyProductService.Describe(uint64(productId))
	if err != nil {
		pc.sendMessage(inputMessage.Chat.ID, "Product not found", "ProductCommander.Edit")
		return
	}

	err = pc.dummyProductService.Update(uint64(productId), *product)
	if err != nil {
		pc.sendMessage(inputMessage.Chat.ID, "Failed to update product", "ProductCommander.Edit")
		return
	}

	pc.sendMessage(inputMessage.Chat.ID, "Product updated successfully", "ProductCommander.Edit")
}
