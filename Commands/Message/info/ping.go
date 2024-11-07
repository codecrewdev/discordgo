// Commands/info/ping.go
package info

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func PingCommand(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	latency := s.HeartbeatLatency().Milliseconds()
	response := fmt.Sprintf("퐁! %dms", latency)
	_, _ = s.ChannelMessageSend(m.ChannelID, response)
}