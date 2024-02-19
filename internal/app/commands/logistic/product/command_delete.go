package product

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
)

func (pc *DummyProductCommander) Delete(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	productId, err := strconv.Atoi(args)
	if err != nil {
		pc.sendMessage(inputMessage.Chat.ID, "Invalid product id", "ProductCommander.Delete")
		return
	}

	_, err = pc.dummyProductService.Remove(uint64(productId))
	if err != nil {
		pc.sendMessage(inputMessage.Chat.ID, "Product not found", "ProductCommander.Delete")
		return
	}

	pc.sendMessage(inputMessage.Chat.ID, "Product successfully deleted", "ProductCommander.Delete")
}
