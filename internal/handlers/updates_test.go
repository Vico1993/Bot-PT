package handlers

import (
	"strings"
	"testing"

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
