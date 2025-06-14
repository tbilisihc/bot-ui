package discord_bot

import (
	"fmt"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
)

var BotToken string = os.Getenv("DISCORD_TOKEN")

func checkNilErr(e error) {
 if e != nil {
  log.Fatal("Error message")
 }
}

func Send(message string) {
	fmt.Println("Discord sending...")
	discord, err := discordgo.New("Bot " + BotToken)
	checkNilErr(err)
	discord.Open()
	fmt.Println(message)
	fmt.Println (os.Getenv("DISCORD_TOKEN"))
	discord.ChannelMessageSend("1383195767623127111", message)
	fmt.Println("Message sent")
	discord.Close()

}