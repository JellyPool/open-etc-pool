package telegram

import (
	"fmt"
	"os"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Telegram configs
var bot, _ = tgbotapi.NewBotAPI(os.Getenv("TelegramBotToken"))

func SendMessage(message string) {
	fmt.Println("TOKEN:" ,os.Getenv("TelegramBotToken"))
	msg := tgbotapi.NewMessage(-1001601201441, message)
	bot.Send(msg)
}