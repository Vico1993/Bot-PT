package handlers

import (
	"fmt"

	"github.com/Vico1993/Bot-PT/internal/utils"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	validType = []string{"group", "supergroup"}
)

func shouldAct(update tgbotapi.Update) bool {
	// If it's not a Group chat
	// If it's not a Message or CallBackQuery
	return utils.InSlice(update.FromChat().Type, validType) &&
		!(update.Message == nil && update.CallbackQuery == nil)
}

func HandleUpdates(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	if !shouldAct(update) {
		fmt.Println("No way")
		// TODO: Log this why we don't act upon it
		return
	}

	msg := tgbotapi.NewMessage(
		update.Message.Chat.ID,
		"I'm listening",
	)

	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}
}
