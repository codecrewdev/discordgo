package utils

import (
	"fmt"
	"runtime"
	"time"

	"github.com/bwmarrin/discordgo"
)

func BotInfoCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// Defer initial response to acknowledge interaction
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
	})
	if err != nil {
		fmt.Println("응답을 연기하는 동안 오류가 발생했습니다:", err)
		return
	}

	developerID := "534214957110394881"
	developerUser, err := s.User(developerID)
	if err != nil {
		fmt.Println("개발자 사용자를 가져오는 동안 오류가 발생했습니다:", err)
		return
	}

	botUser, err := s.User(s.State.User.ID)
	if err != nil {
		fmt.Println("봇 사용자를 가져오는 동안 오류가 발생했습니다:", err)
		return
	}

	shardID := s.ShardID
	botEmbed := &discordgo.MessageEmbed{
		Title:       "봇 정보",
		Description: fmt.Sprintf("> 서버 수: %d\n> 총유저 수: %d\n\u2514유저: %d명 | 봇: %d개\n\n> 개발자: %s\n\n> Golang 버전: %s\n> discordgo 버전: %s\n\n샤드 #%d",
			len(s.State.Guilds),
			countTotalMembers(s),
			countMembers(s, false),
			countMembers(s, true),
			developerUser.Username,
			runtime.Version()[2:],
			discordgo.VERSION,
			shardID),
		Color:     0x00ffcc,
		Thumbnail: &discordgo.MessageEmbedThumbnail{URL: botUser.AvatarURL("1024")},
	}

	time.Sleep(1 * time.Second) // Sleep to ensure interaction is ready for editing

	_, err = s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
		Embeds: &[]*discordgo.MessageEmbed{botEmbed},
	})
	if err != nil {
		fmt.Println("오류 발생: 메시지를 보낼 수 없습니다.", err)
	}
}

func countTotalMembers(s *discordgo.Session) int {
	totalCount := 0
	for _, guild := range s.State.Guilds {
		members, err := s.GuildMembers(guild.ID, "", 1000)
		if err != nil {
			fmt.Println("길드 회원을 가져오는 동안 오류가 발생했습니다:", guild.ID, err)
			continue
		}
		totalCount += len(members)
	}
	return totalCount
}

func countMembers(s *discordgo.Session, isBot bool) int {
	count := 0
	for _, guild := range s.State.Guilds {
		members, err := s.GuildMembers(guild.ID, "", 1000)
		if err != nil {
			fmt.Println("길드 회원을 가져오는 동안 오류가 발생했습니다:", guild.ID, err)
			continue
		}
		for _, member := range members {
			if member.User.Bot == isBot {
				count++
			}
		}
	}
	return count
}

