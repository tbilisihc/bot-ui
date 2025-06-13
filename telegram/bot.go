package telegram_bot

import (
	"context"
	"fmt"
	"os"

	"github.com/go-telegram/bot"
)
func Send(message string) {
	fmt.Println("Telegram sending...")
	var channel_ids = []int64{
		-1002556120690, // Particularly my channels i want to manage
		-1002158048191,
		-1002351566952,
	}
	

	opts := []bot.Option{
	
	}

	b, err := bot.New(os.Getenv("BOT_TOKEN"), opts...)
	if err != nil {
		panic(err)
	}
	fmt.Println(message, "connected to bot")

	for _, chat_id := range channel_ids {
		fmt.Printf("sending to %d\n", chat_id)
		b.SendMessage(context.Background(), &bot.SendMessageParams{
			ChatID: chat_id,
			Text:   message,
		})
	}

	fmt.Println("Message sent")
}

