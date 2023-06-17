package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func SendReply(bot *tgbotapi.BotAPI, chatID int64, messageID int, replyText string) {
	msg := tgbotapi.NewMessage(chatID, replyText)
	msg.ReplyToMessageID = messageID
	_, err := bot.Send(msg)
	if err != nil {
		log.Println("Failed to send reply:", err)
	}
}
