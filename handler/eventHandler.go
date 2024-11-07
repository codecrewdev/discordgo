// handler/eventHandler.go
package handler

import (
	"github.com/bwmarrin/discordgo"
	"github.com/codecrewdev/discordgo/Events/Client" // Adjust the import path as needed
	"github.com/codecrewdev/discordgo/Events/Guilds" // Adjust the import path as needed
	"github.com/codecrewdev/discordgo/Commands/SlashCommands/Games"
)

func RegisterEventHandlers(dg *discordgo.Session) {
	dg.AddHandler(client.Ready)
	dg.AddHandler(guilds.InteractionCreate)
	dg.AddHandler(guilds.MessageCreate)
	dg.AddHandler(client.OnGuildJoin) 
	dg.AddHandler(client.OnGuildRemove) 
	// dg.AddHandler(game_button.ButtonHandler)
	dg.AddHandler(games.ButtonHandler)
}
