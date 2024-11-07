package games

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/codecrewdev/discordgo/model/gamedb"
	"strconv"
	"strings"
)

func formatMoney(money int) string {
	moneyStr := strconv.FormatInt(int64(money), 10)
	n := len(moneyStr)
	if n <= 3 {
		return moneyStr
	}
	var sb strings.Builder
	for i, digit := range moneyStr {
		if i > 0 && (n-i)%3 == 0 {
			sb.WriteRune(',')
		}
		sb.WriteRune(digit)
	}
	return sb.String()
}

func RpdlawjdCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
	})
	if err != nil {
		fmt.Println("응답을 연기하는 동안 오류가 발생했습니다:", err)
		return
	}

	// 유저명 옵션 가져오기
	var user *discordgo.User // 사용자 정보를 저장할 변수를 전역으로 선언
	options := i.Interaction.ApplicationCommandData().Options
	for _, option := range options {
		if option.Name == "유저명" && option.Type == discordgo.ApplicationCommandOptionUser {
			// UserValue를 호출하여 사용자 정보 가져오기
			user = option.UserValue(s) // UserValue 함수 호출
		}
	}

	if user == nil {
		if i.Member != nil {
			user = i.Member.User
		} else {
			user = i.User
		}
	}

	// VerifyUser를 호출하여 유저 정보 가져오기
	isRegistered, money, joinTime, err := gamedb.VerifyUser(user.ID)
	if err != nil {
		fmt.Println("유저 정보를 가져오는 동안 오류가 발생했습니다:", err)
		return
	}

	if !isRegistered {
		Embed := &discordgo.MessageEmbed{
			Title:       "오류",
			Description: fmt.Sprintf("%s 님은 미 가입 사용자입니다.\n\n-# `/가입`으로 가입해주세요.",user.GlobalName),
			Color:       0xFF0000,
			
		}
		_, err = s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
			Embeds: &[]*discordgo.MessageEmbed{Embed},
		})
		if err != nil {
			fmt.Println("오류 발생: 메시지를 보낼 수 없습니다.", err)
		}
		return
	} 

	// Embed 생성 및 유저명 추가
	Embed := &discordgo.MessageEmbed{
		Title: fmt.Sprintf("게임정보 - %s", user.GlobalName),
		Color: 0x00ffcc,
		Thumbnail: &discordgo.MessageEmbedThumbnail{URL: user.AvatarURL("1024")},
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "가입 일",
				Value:  fmt.Sprintf("<t:%d:F> (<t:%d:R>)", joinTime.Unix(), joinTime.Unix()),
				Inline: false,
			},
			{
				Name:   "머니",
				Value:  fmt.Sprintf("%s 원", formatMoney(money)),
				Inline: false,
			},
		},
	}


	_, err = s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
		Embeds: &[]*discordgo.MessageEmbed{Embed},
	})
	if err != nil {
		fmt.Println("오류 발생: 메시지를 보낼 수 없습니다.", err)
	}
}