package games

import (
	"fmt"
	"time"

	"github.com/codecrewdev/discordgo/model/gamedb"
	"github.com/bwmarrin/discordgo"
)

// 명령어 실행 시 사용자의 ID를 저장할 변수
var AuthorizedUserID string

// AccessionCommand는 가입 프로세스를 처리하는 함수입니다.
func AccessionCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// 명령어 실행자의 ID를 설정
	AuthorizedUserID = i.Member.User.ID

	userID := i.Member.User.ID

	// 사용자가 이미 가입했는지 확인
	isRegistered, err := gamedb.IsUserRegistered(userID)
	if err != nil {
		fmt.Println("오류 발생: 사용자 확인 중 오류가 발생했습니다.", err)
		return
	}

	if isRegistered {
		Embed := &discordgo.MessageEmbed{
			Title:       "오류",
			Description: "이미 가입한 사용자입니다..\n\n-# `/탈퇴`으로 탈퇴하실 수 있어요.",
			Color:       0xFF0000,
		}
		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{Embed},
				Flags:  discordgo.MessageFlagsEphemeral,
			},
		})
		if err != nil {
			fmt.Println("오류 발생: 메시지를 보내는 동안 오류가 발생했습니다.", err)
		}
		return
	}

	// 동의 및 미동의 버튼 생성
	buttons := []discordgo.MessageComponent{
		discordgo.Button{
			Label:    "동의",
			Style:    discordgo.SuccessButton,
			CustomID: "agree_button",
		},
		discordgo.Button{
			Label:    "미동의",
			Style:    discordgo.DangerButton,
			CustomID: "disagree_button",
		},
	}

	// 버튼을 포함한 액션 행 생성
	actionRow := discordgo.ActionsRow{
		Components: buttons,
	}

	// 버튼을 포함한 메시지 전송
	Embed := &discordgo.MessageEmbed{
		Title:       "가입하기",
		Description: "[서비스 이용약관]\n[개인정보처리방침]\n\n-# 위 내용 동의하면 아래있는 `동의버튼` 클릭해주세요!",
		Color:       0x00ffcc,
	}

	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds:    []*discordgo.MessageEmbed{Embed},
			Components: []discordgo.MessageComponent{actionRow},
		},
	})
	if err != nil {
		fmt.Println("오류 발생: 메시지를 보내는 동안 오류가 발생했습니다.", err)
	}

	// 버튼 비활성화 타이머 실행
	go disableButtonsAfterTimeout(s, i, 30*time.Second)
}

// 버튼을 30초 후 비활성화하는 함수
func disableButtonsAfterTimeout(s *discordgo.Session, i *discordgo.InteractionCreate, duration time.Duration) {
	time.Sleep(duration)

	// 비활성화된 버튼 생성
	disabledButtons := []discordgo.MessageComponent{
		discordgo.Button{
			Label:    "동의",
			Style:    discordgo.SuccessButton,
			CustomID: "agree_button",
			Disabled: true,
		},
		discordgo.Button{
			Label:    "미동의",
			Style:    discordgo.DangerButton,
			CustomID: "disagree_button",
			Disabled: true,
		},
	}

	// 비활성화된 버튼을 포함한 액션 행 생성
	disabledActionRow := discordgo.ActionsRow{
		Components: disabledButtons,
	}

	Embed := &discordgo.MessageEmbed{
		Title:       "가입하기",
		Description: "[서비스 이용약관]\n[개인정보처리방침]\n\n-# 위 내용 동의하면 아래있는 `동의버튼` 클릭해주세요!",
		Color:       0x00ffcc,
	}

	// 메시지를 업데이트하여 버튼을 비활성화
	_, err := s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
		Embeds:    &[]*discordgo.MessageEmbed{Embed},
		Components: &[]discordgo.MessageComponent{disabledActionRow},
	})
	if err != nil {
		fmt.Println("오류 발생: 버튼 비활성화 업데이트 중 오류가 발생했습니다.", err)
	}
}

// ButtonHandler는 버튼 클릭 이벤트를 처리하는 함수입니다.
func ButtonHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// Interaction의 타입이 버튼 클릭인지 확인
	if i.Type != discordgo.InteractionMessageComponent {
		return
	}

	// 명령어 실행자가 아닌 사용자는 버튼 사용 불가
	if i.Member.User.ID != AuthorizedUserID {
		errEmbed := &discordgo.MessageEmbed{
			Title:       "접근 불가",
			Description: "이 버튼은 명령어를 실행한 사용자만 사용할 수 있습니다.",
			Color:       0xFF0000,
		}
		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{errEmbed},
				Flags:  discordgo.MessageFlagsEphemeral,
			},
		})
		if err != nil {
			fmt.Println("오류 발생: 접근 불가 응답 처리 중 오류가 발생했습니다.", err)
		}
		return
	}

	// 버튼이 "agree_button" 또는 "disagree_button"인지 확인 후 처리
	switch i.MessageComponentData().CustomID {
	case "agree_button":
		// 동의 버튼 클릭 시 처리
		userID := i.Member.User.ID
		joinTime := time.Now()
		money := 5000 // 초기 머니 값 설정
		gamedb.Accessiondb(userID, money, joinTime)

		agreeEmbed := &discordgo.MessageEmbed{
			Title:       "가입 완료",
			Description: "이용약관 및 개인정보처리방침에 동의하셨습니다.\n이제 서비스를 이용하실 수 있습니다.",
			Color:       0x00FF00,
		}
		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseUpdateMessage,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{agreeEmbed},
			},
		})
		if err != nil {
			fmt.Println("오류 발생: 동의 버튼 응답 처리 중 오류가 발생했습니다.", err)
		}

	case "disagree_button":
		// 미동의 버튼 클릭 시 처리
		disagreeEmbed := &discordgo.MessageEmbed{
			Title:       "가입 취소",
			Description: "이용약관 및 개인정보처리방침에 동의하지 않으셨습니다.\n가입이 취소되었습니다.",
			Color:       0xFF0000,
		}
		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseUpdateMessage,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{disagreeEmbed},
			},
		})
		if err != nil {
			fmt.Println("오류 발생: 미동의 버튼 응답 처리 중 오류가 발생했습니다.", err)
		}
	}
}
