// handler/commandHandler.go
package handler

import (
	"yourbot/Commands/info"        // Adjust this path as necessary
	"yourbot/SlashCommands/Utils"   // Adjust this path as necessary
	"yourbot/Events/Guilds"         // Adjust this path as necessary
)

func RegisterCommands() {
	// Register text commands
	guilds.TextCommands["핑"] = info.PingCommand

	// Register slash commands
	guilds.SlashCommands["핑"] = utils.PingSlashCommand
	guilds.SlashCommands["개발자"] = utils.DeveloperCommand
	guilds.SlashCommands["봇정보"] = utils.BotInfoCommand
}