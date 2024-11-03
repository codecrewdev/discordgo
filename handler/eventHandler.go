// handler/eventHandler.go
package handler

import (
	"github.com/bwmarrin/discordgo"
	"yourbot/Events/Client" // Adjust the import path as needed
	"yourbot/Events/Guilds" // Adjust the import path as needed
)

func RegisterEventHandlers(dg *discordgo.Session) {
	dg.AddHandler(client.Ready)
	dg.AddHandler(guilds.InteractionCreate)
	dg.AddHandler(guilds.MessageCreate)
}
