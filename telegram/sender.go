package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Telegram configs
const telegram_token string = ""
var bot, _ = tgbotapi.NewBotAPI(telegram_token)

func SendMessage(message string) {
	msg := tgbotapi.NewMessage(-1001601201441, message)
	bot.Send(msg)
}