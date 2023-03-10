package handlers

import (
	"strings"
	"testing"

	"github.com/Vico1993/Bot-PT/internal/chatgpt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/stretchr/testify/assert"
)

func TestShouldActNotActBecauseNotInGroup(t *testing.T) {
	res := shouldAct(tgbotapi.Update{
		Message: &tgbotapi.Message{
			Chat: &tgbotapi.Chat{
				Type: "private",
			},
		},
	})

	assert.False(t, res, "Should be false as the type is private not "+strings.Join(validType, ""))
}

func TestShouldActNotActBecauseNotMessageOrCallBack(t *testing.T) {
	res := shouldAct(tgbotapi.Update{
		CallbackQuery: nil,
		Message:       nil,
		EditedMessage: &tgbotapi.Message{
			Chat: &tgbotapi.Chat{
				Type: "private",
			},
		},
	})

	assert.False(t, res, "Should be false as Message and CallbackQuery are nil")
}

func TestExtractResponseTxtFromAChatResponse(t *testing.T) {
	response := chatgpt.Response{
		Choices: []chatgpt.Choice{
			{
				Message: chatgpt.Message{
					Role:    "bot",
					Content: "Hello",
				},
			},
		},
	}

	result := extractResponseTxt(&response)

	assert.Equal(t, "Hello", result, "Should return Hello")
}

func TestExtractResponseTxtEmptyChoice(t *testing.T) {
	response := chatgpt.Response{
		Choices: []chatgpt.Choice{},
	}

	result := extractResponseTxt(&response)

	assert.Equal(t, "Not sure what to answer", result, "Should return the default message if no choice returned from API")
}

func TestExtractResponseTxtNil(t *testing.T) {
	result := extractResponseTxt(nil)

	assert.Equal(t, "Not sure what to answer", result, "Should return the default message if no response returned")
}

func TestBuildMessage(t *testing.T) {
	res := buildMessage(1234, 12, "Hello")

	assert.Equal(t, "HTML", res.ParseMode, "Parse Mode should be HTML")
	assert.Equal(t, 12, res.ReplyToMessageID, "Reply Message should be correctly set")
	assert.Equal(t, int64(1234), res.ChatID, "Message ID should be correctly set")
	assert.Equal(t, "Hello", res.Text, "Text should be correctly set")
}
