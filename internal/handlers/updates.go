package handlers

import (
	"fmt"
	"strconv"

	"github.com/Vico1993/Bot-PT/internal/chatgpt"
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

func extractResponseTxt(chatgptResponse *chatgpt.Response) string {
	if chatgptResponse == nil || len(chatgptResponse.Choices) == 0 {
		return "Not sure what to answer"
	}

	return chatgptResponse.Choices[0].Message.Content
}

func buildMessage(chatID int64, replyID int, text string) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(
		chatID,
		text,
	)
	msg.ReplyToMessageID = replyID
	msg.ParseMode = tgbotapi.ModeHTML

	return msg
}

func HandleUpdates(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	if !shouldAct(update) {
		// TODO: Log this why we don't act upon it
		return
	}

	response := chatgpt.Ask(update.Message.Text)

	_, err := bot.Send(
		buildMessage(
			update.Message.Chat.ID,
			update.Message.MessageID,
			extractResponseTxt(response),
		),
	)

	if err != nil {
		fmt.Println("Couldn't send message")
	}

	if response != nil {
		_, err := bot.Send(
			buildMessage(
				update.Message.Chat.ID,
				update.Message.MessageID,
				"This message cost you: "+strconv.Itoa(int(response.Usage.TotalTokens))+" tokens",
			),
		)

		if err != nil {
			fmt.Println("Couldn't send message")
		}
	}
}
