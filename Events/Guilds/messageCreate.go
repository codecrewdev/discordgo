// Events/Guilds/messageCreate.go
package guilds

import (
	"log"
	"strings"
	"github.com/bwmarrin/discordgo"
	"github.com/codecrewdev/discordgo/Commands/Message/info" // Adjust the import path as needed
)

const prefix = "!"

var TextCommands = map[string]func(*discordgo.Session, *discordgo.MessageCreate, []string){
	"핑": info.PingCommand,
	"ping": info.PingCommand,
}

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot || !strings.HasPrefix(m.Content, prefix) {
		return
	}

	args := strings.Fields(m.Content[len(prefix):])
	commandName := strings.ToLower(args[0])
	guildName := getGuildName(s, m.GuildID)

	if command, exists := TextCommands[commandName]; exists {
		command(s, m, args[1:])
		log.Printf("명령 사용됨: %s by %s (%s) in guild %s (%s)", commandName, m.Author.Username, m.Author.ID, guildName, m.GuildID)
	} else {
		log.Printf("\033[31m알 수 없는 명령: %s by %s (%s) in guild %s (%s)\033[0m", commandName, m.Author.Username, m.Author.ID, guildName, m.GuildID)
	}
}

