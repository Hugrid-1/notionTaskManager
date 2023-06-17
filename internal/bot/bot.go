package bot

import (
	"log"

	"github.com/Hugrid-1/notionTaskManager/internal/platform/telegram"
)

func StartBot() error {
	telegramBot, err := telegram.NewBot("6204371684:AAGXIYJtHx7vh3fJHv-me4aLOUokxRM0jUM")
	if err != nil {
		return err
	}

	errCh := make(chan error)

	go func() {
		errCh <- telegramBot.Start()
	}()

	// Handle errors from the Telegram bot
	if err := <-errCh; err != nil {
		log.Println("Bot error:", err)
	}

	return nil
}
