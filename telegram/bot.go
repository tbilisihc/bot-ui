package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/go-telegram/bot"
)
func Send(message string) {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
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
	b.Start(ctx)
	for _, chat_id := range channel_ids {

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: chat_id,
		Text:   message,
	})
}
}

