package handler

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

var desiredCommands = []*discordgo.ApplicationCommand{
	{
		Name:        "핑",
		Description: "핑을 보여줍니다",
	},
	{
		Name:        "개발자",
		Description: "개발자 정보 입니다.",
	},
	{
		Name:        "봇정보",
		Description: "현재 봇 정보 입니다.",
	},
	{
		Name:        "가입",
		Description: "가입 통하여 재미있는 게임 해보세요.",
	},
}

func RegisterSlashCommands(dg *discordgo.Session) {
	if dg == nil {
		fmt.Println("오류: Discord 세션이 없습니다.")
		return
	}

	// Register the desired commands
	for _, cmd := range desiredCommands {
		_, err := dg.ApplicationCommandCreate(dg.State.User.ID, "", cmd)
		if err != nil {
			fmt.Printf("'%s' 명령을 생성할 수 없습니다. %v\n", cmd.Name, err)
		}
	}

	// Fetch all currently registered commands
	existingCommands, err := dg.ApplicationCommands(dg.State.User.ID, "")
	if err != nil {
		fmt.Println("애플리케이션 명령을 가져오는 중 오류가 발생했습니다:", err)
		return
	}

	// Convert desired commands to a map for easy lookup
	desiredCommandMap := make(map[string]bool)
	for _, cmd := range desiredCommands {
		desiredCommandMap[cmd.Name] = true
	}

	// Delete commands that are not in the desired list
	for _, cmd := range existingCommands {
		if !desiredCommandMap[cmd.Name] {
			err := dg.ApplicationCommandDelete(dg.State.User.ID, "", cmd.ID)
			if err != nil {
				fmt.Printf("명령 '%s' 삭제 중 오류: %v\n", cmd.Name, err)
			} else {
				fmt.Printf("불필요한 명령을 삭제했습니다: %s\n", cmd.Name)
			}
		}
	}
}