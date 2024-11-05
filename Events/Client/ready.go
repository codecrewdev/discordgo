package client

import (
    "log"
    "fmt"
    "sync"
    "github.com/bwmarrin/discordgo"
)

var (
    initialGuilds   = make(map[string]bool)
    initialGuildsMu sync.Mutex
)

// Ready 이벤트 핸들러
func Ready(s *discordgo.Session, r *discordgo.Ready) {
    initialGuildsMu.Lock()
    defer initialGuildsMu.Unlock()

    // 봇이 처음 시작될 때 현재 가입된 모든 길드를 저장
    for _, guild := range r.Guilds {
        initialGuilds[guild.ID] = true
    }

    fmt.Printf("%s 로 로그인되었습니다.\n", s.State.User.Username)
}

// OnGuildJoin 이벤트 핸들러
func OnGuildJoin(s *discordgo.Session, guildCreate *discordgo.GuildCreate) {
    initialGuildsMu.Lock()
    defer initialGuildsMu.Unlock()

    // 초기 길드 목록에 존재하는 경우라면 처리하지 않음
    if _, exists := initialGuilds[guildCreate.Guild.ID]; exists {
        // 초기 길드 목록에서 제거 (처음 이후에는 새로운 초대를 처리하기 위해)
        delete(initialGuilds, guildCreate.Guild.ID)
        return
    }

    // 새로운 길드에 초대된 경우 처리
    owner := guildCreate.Guild.OwnerID
    ownerUser, err := s.User(owner)
    if err != nil {
        log.Println("오너 정보를 가져오는 데 오류 발생:", err)
        return
    }

    logMessage := fmt.Sprintf("%s(%s)에 봇 초대됨. owner = %s(%s) (서버 수: %d개)\n",
        guildCreate.Guild.Name, guildCreate.Guild.ID, ownerUser.Username, ownerUser.ID, len(s.State.Guilds))

    log.Print(logMessage)
}

// OnGuildRemove 이벤트 핸들러
func OnGuildRemove(s *discordgo.Session, guildDelete *discordgo.GuildDelete) {
    guild := guildDelete.Guild
    if guild == nil {
        log.Println("길드 정보가 없습니다.")
        return
    }

    owner := guild.OwnerID
    ownerUser, err := s.User(owner)
    if err != nil {
        log.Println("오너 정보를 가져오는 데 오류 발생:", err)
        return
    }

    logMessage := fmt.Sprintf("%s(%s)에 봇 추방됨. owner = %s(%s) (서버 수: %d개)\n",
        guild.Name, guild.ID, ownerUser.Username, ownerUser.ID, len(s.State.Guilds))

    log.Print(logMessage)
}
