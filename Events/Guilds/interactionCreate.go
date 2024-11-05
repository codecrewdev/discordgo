package guilds

import (
	"log"
	"github.com/bwmarrin/discordgo"
	"yourbot/SlashCommands/Utils" // Adjust the import path as needed
	"yourbot/SlashCommands/information"   // Adjust this path as necessary
)

// Register slash commands in a map
var SlashCommands = map[string]func(*discordgo.Session, *discordgo.InteractionCreate){
	"핑": utils.PingSlashCommand,
	"개발자" : information.DeveloperCommand,
	"봇정보" : information.BotInfoCommand,
}

// InteractionCreate handles slash commands
func InteractionCreate(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type != discordgo.InteractionApplicationCommand {
		return
	}

	commandName := i.ApplicationCommandData().Name
	guildName := getGuildName(s, i.GuildID)
	user := i.Member.User

	if command, exists := SlashCommands[commandName]; exists {
		command(s, i)
		log.Printf("슬래시 명령 사용됨: %s by %s (%s) in guild %s (%s)", commandName, user.Username, user.ID, guildName, i.GuildID)
	} else {
		log.Printf("\033[31m알 수 없는 슬래시 명령: %s by %s (%s) in guild %s (%s)\033[0m", commandName, user.Username, user.ID, guildName, i.GuildID)
	}
}

// getGuildName retrieves the name of the guild by its ID
func getGuildName(s *discordgo.Session, guildID string) string {
	guild, err := s.State.Guild(guildID)
	if err != nil || guild == nil {
		return "Unknown Guild"
	}
	return guild.Name
}