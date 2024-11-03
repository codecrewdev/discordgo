// Events/Guilds/interactionCreate.go
package guilds

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"yourbot/SlashCommands/Utils" // Adjust the import path as needed
)

// Register slash commands in a map
var SlashCommands = map[string]func(*discordgo.Session, *discordgo.InteractionCreate){
	"핑": utils.PingSlashCommand,
}

// InteractionCreate handles slash commands
func InteractionCreate(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type != discordgo.InteractionApplicationCommand {
		return
	}

	if command, exists := SlashCommands[i.ApplicationCommandData().Name]; exists {
		command(s, i)
	} else {
		fmt.Println("Unknown slash command:", i.ApplicationCommandData().Name)
	}
}
