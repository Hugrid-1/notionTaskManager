package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"

	"github.com/Hugrid-1/notionTaskManager/internal/bot"
)

type Bot struct {
	bot *tgbotapi.BotAPI
}

func NewBot(token string) (*Bot, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	bot.Debug = true

	return &Bot{bot: bot}, nil
}

func (b *Bot) Start() error {
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	updates, err := b.bot.GetUpdatesChan(updateConfig)
	if err != nil {
		return err
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		message := update.Message

		if update.Message.IsCommand() {
			switch update.Message.Command() {
			case "add":
				bot.HandleAddCommand(b, message)
			case "addUser":
				bot.HandleAddUserCommand(b, message)
			default:
				reply := "Unknown command."
				b.SendReply(message.Chat.ID, message.MessageID, reply)
			}
		} else {
			if !bot.IsAllowedUser(message.From.ID) {
				reply := "Please contact the administrator."
				b.SendReply(message.Chat.ID, message.MessageID, reply)
			} else {
				reply := "Record successfully added!"
				b.SendReply(message.Chat.ID, message.MessageID, reply)
				// Add code here to add the record to Notion using the notionClient
			}
		}
	}

	return nil
}

func (b *Bot) SendReply(chatID int64, messageID int, replyText string) {
	msg := tgbotapi.NewMessage(chatID, replyText)
	msg.ReplyToMessageID = messageID
	_, err := b.bot.Send(msg)
	if err != nil {
		log.Println("Failed to send reply:", err)
	}
}
