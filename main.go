// main.go
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"yourbot/handler" // Adjust the import path as needed

	"github.com/joho/godotenv"
	"github.com/bwmarrin/discordgo"
)




func main() {

	// .env 파일 로드
    err := godotenv.Load()
    if err != nil {
        fmt.Println("Error loading .env file")
        return
    }

    // 환경 변수에서 토큰 가져오기
    token := os.Getenv("TOKEN")
    if token == "" {
        fmt.Println("Token is not set in .env file")
        return
    }
	// Set up the intents
	intents := discordgo.IntentsGuilds |
		discordgo.IntentsGuildMessages |
		discordgo.IntentsGuildMessageReactions |
		discordgo.IntentsGuildPresences | // Presence Intent
		discordgo.IntentsGuildMembers |   // Server Members Intent
		discordgo.IntentsMessageContent   // Message Content Intent

	// Create a new Discord session with the specified intents and shard configuration
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Error creating Discord session:", err)
		return
	}
	dg.Identify.Intents = intents
	dg.Identify.Shard = nil

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


