// Events/Client/ready.go
package client

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func Ready(s *discordgo.Session, r *discordgo.Ready) {
	fmt.Printf("%s 로 로그인되었습니다.\n", s.State.User.Username)
}
