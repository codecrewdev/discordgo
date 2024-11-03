// Events/Guilds/messageCreate.go
package guilds

import (
	"fmt"
	"strings"
	"github.com/bwmarrin/discordgo"
	"yourbot/Commands/info" // Adjust the import path as needed
)

const prefix = "!"

var TextCommands = map[string]func(*discordgo.Session, *discordgo.MessageCreate, []string){
	"핑": info.PingCommand,
}

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot || !strings.HasPrefix(m.Content, prefix) {
		return
	}

	args := strings.Fields(m.Content[len(prefix):])
	commandName := strings.ToLower(args[0])

	if command, exists := TextCommands[commandName]; exists {
		command(s, m, args[1:])
	} else {
		fmt.Println("Unknown command:", commandName)
	}
}
