package main

import (
	"github.com/Hugrid-1/notionTaskManager/internal/bot"
	"log"
)

func main() {
	err := bot.StartBot()
	if err != nil {
		log.Fatal(err)
	}
}
