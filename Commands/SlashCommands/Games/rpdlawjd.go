package games

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func RpdlawjdCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
	})
	if err != nil {
		fmt.Println("응답을 연기하는 동안 오류가 발생했습니다:", err)
		return
	}

	user := i.Member

	Embed:= &discordgo.MessageEmbed{
		Title: fmt.Sprintf("게임정보 - %s", user.User.GlobalName),

	}

	_, err = s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
		Embeds: &[]*discordgo.MessageEmbed{Embed},
	})
	if err != nil {
		fmt.Println("오류 발생: 메시지를 보낼 수 없습니다.", err)
	}
}