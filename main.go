// main.go
package main

import (
	"log"
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
        log.Println(".env 파일 로딩 오류")
        return
    }

    // 환경 변수에서 토큰 가져오기
    token := os.Getenv("TOKEN")
    if token == "" {
        log.Println(".env 파일에 토큰이 설정되지 않았습니다.")
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
		log.Println("Discord 세션 생성 중 오류 발생:", err)
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
		log.Println("연결을 여는 중 오류가 발생했습니다:", err)
		return
	}
	defer func() {
		dg.Close()
		fmt.Println("봇 연결이 종료되었습니다.")
	}()

	// Register slash commands only after the connection is established
	handler.RegisterSlashCommands(dg)

	fmt.Println("봇이 실행 중입니다. 종료하려면 CTRL+C를 누르세요.")
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-stop

	fmt.Println("프로그램이 종료됩니다.")
}


