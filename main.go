package main

import (
	discord_bot "bot-ui/discord"
	get_conf "bot-ui/func"
	telegram_bot "bot-ui/telegram"
	"fmt"
)


func main() {
	get_conf.Read()
	var message string
	fmt.Print("Hello, to sent a message to your telegram channels and discord type it here: ")
	fmt.Scan(&message)
	telegram_bot.Send(message)
	discord_bot.Send(message)
}