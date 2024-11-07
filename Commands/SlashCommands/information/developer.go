package information

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func DeveloperCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
	})
	if err != nil {
		fmt.Println("응답을 연기하는 동안 오류가 발생했습니다:", err)
		return
	}
	
	// 개발자 정보를 가져옵니다.
	developerID := "534214957110394881"
	developerUser, err := s.User(developerID)
	if err != nil {
		fmt.Println("오류 발생: 개발자 정보를 가져올 수 없습니다.", err)
		return
	}

	// 임베드 메시지 생성
	embed := &discordgo.MessageEmbed{
		Title:       "개발자 정보",
		Description: fmt.Sprintf("- %s (%s)", developerUser.GlobalName, developerUser.Username),
		Color:       0x00ffcc, // 사용자 색상 (고정된 값 사용)
	}

	// 메시지 보내기
	_, err = s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
		Embeds: &[]*discordgo.MessageEmbed{embed},
	})
	if err != nil {
		fmt.Println("오류 발생: 메시지를 보낼 수 없습니다.", err)
	}
}