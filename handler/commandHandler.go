// handler/commandHandler.go
package handler

import (
	"github.com/codecrewdev/discordgo/Commands/Message/info" // Adjust this path as necessary
	"github.com/codecrewdev/discordgo/Events/Guilds" // Adjust this path as necessary
	"github.com/codecrewdev/discordgo/Commands/SlashCommands/Games"
	"github.com/codecrewdev/discordgo/Commands/SlashCommands/Utils"       // Adjust this path as necessary
	"github.com/codecrewdev/discordgo/Commands/SlashCommands/information" // Adjust this path as necessary
)

func RegisterCommands() {
	// Register text commands
	guilds.TextCommands["핑"] = info.PingCommand
	guilds.TextCommands["ping"] = info.PingCommand

	// Register slash commands
	guilds.SlashCommands["핑"] = utils.PingSlashCommand
	guilds.SlashCommands["개발자"] = information.DeveloperCommand
	guilds.SlashCommands["봇정보"] = information.BotInfoCommand
	guilds.SlashCommands["가입"] = games.AccessionCommand
	guilds.SlashCommands["게임정보"] = games.RpdlawjdCommand
}