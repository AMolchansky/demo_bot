package product

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
)

func (pc *DummyProductCommander) Get(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	productId, err := strconv.Atoi(args)
	if err != nil {
		pc.sendMessage(inputMessage.Chat.ID, "Invalid product id", "ProductCommander.Get")
		return
	}

	product, err := pc.dummyProductService.Describe(uint64(productId))
	if err != nil {
		pc.sendMessage(inputMessage.Chat.ID, "Product not found", "ProductCommander.Get")
		return
	}

	pc.sendMessage(inputMessage.Chat.ID, product.Title, "ProductCommander.Get")
}
