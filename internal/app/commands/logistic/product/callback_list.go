package product

import (
	"encoding/json"
	"fmt"
	"github.com/AMolchansky/demo_bot/internal/app/path"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strings"
)

type CallbackListData struct {
	Offset int `json:"offset"`
	Page   int `json:"page"`
}

func (pc *DummyProductCommander) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	outputMsg := strings.Builder{}
	parsedData := CallbackListData{}

	err := json.Unmarshal([]byte(callbackPath.CallbackData), &parsedData)
	if err != nil {
		log.Printf("ProductCommander.CallbackList: "+
			"error reading json data for type CallbackListData from "+
			"input string %v - %v", callbackPath.CallbackData, err)
		return
	}

	products, errIndex := pc.dummyProductService.List(uint64(parsedData.Page), uint64(parsedData.Offset))

	if errIndex != nil {
		pc.sendMessage(callback.Message.Chat.ID, "Invalid page", "ProductCommander.CallbackList")
		return
	}

	if len(products) == 0 {
		pc.sendMessage(callback.Message.Chat.ID, "Products not found", "ProductCommander.CallbackList")
		return
	}

	for _, p := range products {
		outputMsg.WriteString(fmt.Sprintf("[ID: %d] %s \n", p.Id, p.Title))
	}

	msg := tgbotapi.NewEditMessageText(callback.Message.Chat.ID, callback.Message.MessageID, outputMsg.String())

	pc.addPrevButton(&msg, parsedData.Page, parsedData.Offset)
	pc.addNextButton(&msg, parsedData.Page, parsedData.Offset)

	_, err = pc.bot.Send(msg)
	if err != nil {
		log.Printf("ProductCommander.CallbackList: error sending reply message to chat - %v", err)
	}
}

func (pc *DummyProductCommander) addPrevButton(msg *tgbotapi.EditMessageTextConfig, page, offset int) {
	if page <= 1 {
		return
	}
	pc.addButton(msg, page-1, offset, "Prev page")
}

func (pc *DummyProductCommander) addNextButton(msg *tgbotapi.EditMessageTextConfig, page, offset int) {
	_, err := pc.dummyProductService.List(uint64(page+1), uint64(offset))
	if err == nil {
		pc.addButton(msg, page+1, offset, "Next page")
	}
}

func (pc *DummyProductCommander) addButton(msg *tgbotapi.EditMessageTextConfig, page, offset int, text string) {
	serializedData, err := json.Marshal(CallbackListData{
		Page:   page,
		Offset: offset,
	})

	if err != nil {
		pc.logError("error json.Marshal CallbackListData", err)
		return
	}

	callbackPath := path.CallbackPath{
		Domain:       "logistic",
		Subdomain:    "product",
		CallbackName: "list",
		CallbackData: string(serializedData),
	}

	inlineKeyboardMarkup := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(text, callbackPath.String()),
		),
	)

	msg.ReplyMarkup = &inlineKeyboardMarkup
}

func (pc *DummyProductCommander) logError(context string, err error) {
	log.Printf("ProductCommander: %s - %v", context, err)
}
