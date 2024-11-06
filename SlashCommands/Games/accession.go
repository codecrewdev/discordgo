// SlashCommands/Games/accession.go
package games

import(
	"fmt"
	"time"

	"yourbot/model/gamedb"
	"github.com/bwmarrin/discordgo"
)

func AccessionCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// Defer initial response to acknowledge interaction
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
	})
	if err != nil {
		fmt.Println("응답을 연기하는 동안 오류가 발생했습니다:", err)
		return
	}

	userID := i.Member.User.ID

	// 사용자가 이미 가입했는지 확인
	isRegistered, err := gamedb.VerifyUser(userID)
	if err != nil {
		fmt.Println("오류 발생: 사용자 확인 중 오류가 발생했습니다.", err)
		return
	}

	if isRegistered {
		// 이미 가입한 경우
		message := "이미 가입한 사용자입니다."
		_, err = s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
			Content: &message,
		})
		if err != nil {
			fmt.Println("오류 발생: 메시지를 수정하는 동안 오류가 발생했습니다.", err)
		}
		return
	}

	// 사용자가 가입하지 않은 경우 처리
	joinTime := time.Now()
	money := 0 // 초기 머니 값 설정
	gamedb.Accessiondb(userID, money, joinTime)

	Embed := &discordgo.MessageEmbed{
		Title:       "가입하기",
		Description: "[서비스 이용약관]\n[개인정보처리방침]\n-# 위 내용 동의하면 아래있는 동의버튼 클릭해주세요!",
		Color:       0x00ffcc,
	}
	_, err = s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
		Embeds: &[]*discordgo.MessageEmbed{Embed},
	})
	if err != nil {
		fmt.Println("오류 발생: 메시지를 보낼 수 없습니다.", err)
	}
}
