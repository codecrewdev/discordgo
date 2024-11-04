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
}

func RegisterSlashCommands(dg *discordgo.Session) {
	if dg == nil {
		fmt.Println("Error: Discord session is nil")
		return
	}

	// Register the desired commands
	for _, cmd := range desiredCommands {
		_, err := dg.ApplicationCommandCreate(dg.State.User.ID, "", cmd)
		if err != nil {
			fmt.Printf("Cannot create '%s' command: %v\n", cmd.Name, err)
		}
	}

	// Fetch all currently registered commands
	existingCommands, err := dg.ApplicationCommands(dg.State.User.ID, "")
	if err != nil {
		fmt.Println("Error fetching application commands:", err)
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
				fmt.Printf("Error deleting command '%s': %v\n", cmd.Name, err)
			} else {
				fmt.Printf("Deleted unnecessary command: %s\n", cmd.Name)
			}
		}
	}
}