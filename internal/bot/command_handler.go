package bot

import (
	"strconv"

	"github.com/Hugrid-1/notionTaskManager/internal/platform/telegram"
)

func handleAddCommand(platform interface{}, message interface{}) {
	note := message.Text[5:] // Remove "/add " from the message

	// Add code here to add the record to Notion using the notionClient

	reply := "Record successfully added!"
	telegramBot := platform.(*telegram.Bot)
	telegramBot.SendReply(message.ChatID, message.MessageID, reply)
}

func handleAddUserCommand(platform interface{}, message interface{}) {
	if len(message.CommandArguments()) == 0 {
		reply := "User ID is not specified."
		telegramBot := platform.(*telegram.Bot)
		telegramBot.SendReply(message.ChatID, message.MessageID, reply)
		return
	}

	userID, err := strconv.Atoi(message.CommandArguments())
	if err != nil {
		reply := "Invalid user ID."
		telegramBot := platform.(*telegram.Bot)
		telegramBot.SendReply(message.ChatID, message.MessageID, reply)
		return
	}

	allowedUserIDs = append(allowedUserIDs, userID)

	reply := "User successfully added to the allowed users list."
	telegramBot := platform.(*telegram.Bot)
	telegramBot.SendReply(message.ChatID, message.MessageID, reply)
}
