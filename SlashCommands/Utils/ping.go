// SlashCommands/Utils/ping.go
package utils

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

// pingStatus returns an emoji and description based on latency
func pingStatus(ping int) (string, string) {
	switch {
	case ping <= 100:
		return "🔵", "매우 좋음"
	case ping <= 200:
		return "🟢", "좋음"
	case ping <= 500:
		return "🟡", "보통"
	case ping <= 1000:
		return "🟠", "나쁨"
	default:
		return "🔴", "매우 나쁨"
	}
}

// PingSlashCommand handles the "핑" slash command
func PingSlashCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// Defer initial response
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
	})
	if err != nil {
		fmt.Println("Error deferring response:", err)
		return
	}

	// Measure WebSocket latency
	apiPing := int(s.HeartbeatLatency().Milliseconds())

	// Measure message send latency
	sendStart := time.Now()
	msg, err := s.ChannelMessageSend(i.ChannelID, "메시지핑 측정중 입니다..")
	sendLatency := int(time.Since(sendStart).Milliseconds())
	if err != nil {
		fmt.Println("Error sending message:", err)
		return
	}

	// Measure message edit latency
	editStart := time.Now()
	_, err = s.ChannelMessageEdit(i.ChannelID, msg.ID, "메시지핑 측정중 입니다...")
	editLatency := int(time.Since(editStart).Milliseconds())
	if err != nil {
		fmt.Println("Error editing message:", err)
		return
	}

	// Measure message delete latency
	deleteStart := time.Now()
	err = s.ChannelMessageDelete(i.ChannelID, msg.ID)
	deleteLatency := int(time.Since(deleteStart).Milliseconds())
	if err != nil {
		fmt.Println("Error deleting message:", err)
		return
	}

	// Get emoji and status descriptions for each latency
	apiIcon, apiStatus := pingStatus(apiPing)
	sendIcon, sendStatus := pingStatus(sendLatency)
	editIcon, editStatus := pingStatus(editLatency)
	deleteIcon, deleteStatus := pingStatus(deleteLatency)

	// Build the embed message with latency information
	embed := &discordgo.MessageEmbed{
		Title: "현재 핑",
		Color: 0x3498db, // Set a color for the embed
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: s.State.User.AvatarURL(""),
		},
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   fmt.Sprintf("%s 디스코드 API 핑", apiIcon),
				Value:  fmt.Sprintf("%dms | %s", apiPing, apiStatus),
				Inline: true,
			},
			{
				Name:   fmt.Sprintf("%s 메시지 전송 핑", sendIcon),
				Value:  fmt.Sprintf("%dms | %s", sendLatency, sendStatus),
				Inline: true,
			},
			{
				Name:   fmt.Sprintf("%s 메시지 수정 핑", editIcon),
				Value:  fmt.Sprintf("%dms | %s", editLatency, editStatus),
				Inline: true,
			},
			{
				Name:   fmt.Sprintf("%s 메시지 삭제 핑", deleteIcon),
				Value:  fmt.Sprintf("%dms | %s", deleteLatency, deleteStatus),
				Inline: true,
			},
			{
				Name:   "샤드 수",
				Value:  fmt.Sprintf("%d 개", s.ShardCount),
				Inline: true,
			},
		},
	}

	// Send the response with the embedded message
	embeds := []*discordgo.MessageEmbed{embed} // Define embeds as a slice pointer
	_, err = s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
		Embeds: &embeds,
	})
	if err != nil {
		fmt.Println("Error sending ping response:", err)
	}
}
