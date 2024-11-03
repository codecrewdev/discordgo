// main.go
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"yourbot/handler" // Adjust the import path as needed

	"github.com/bwmarrin/discordgo"
)

var Token = "MTIwNjg3ODY4Njc2NzQ4MDg1Mg.GdBKlM.WmOe91D4a-EWr50VD-4lD5sJwo_WeJ-z36lM1s" // Load from environment or .env if preferred

func main() {
	// Set up the intents
	intents := discordgo.IntentsGuilds |
		discordgo.IntentsGuildMessages |
		discordgo.IntentsGuildMessageReactions |
		discordgo.IntentsGuildPresences | // Presence Intent
		discordgo.IntentsGuildMembers |   // Server Members Intent
		discordgo.IntentsMessageContent   // Message Content Intent

	// Create a new Discord session with the specified intents and shard configuration
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("Error creating Discord session:", err)
		return
	}
	dg.Identify.Intents = intents
	dg.Identify.Shard = &[2]int{0, 1} // Single shard: shard ID 0 out of 1 total shard

	// Register event handlers and commands
	handler.RegisterEventHandlers(dg)
	handler.RegisterCommands()

	// Open a websocket connection to Discord
	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening connection:", err)
		return
	}
	defer dg.Close()

	// Register slash commands only after the connection is established
	handler.RegisterSlashCommands(dg)

	fmt.Println("Bot is now running. Press CTRL+C to exit.")
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-stop
}
