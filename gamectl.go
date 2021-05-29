package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
	embed "github.com/clinet/discordgo-embed"
)

func startGame(m *discordgo.MessageCreate) {
	key := m.GuildID + m.ChannelID
	g := newGame(m.GuildID, m.ChannelID, m.Author.ID)
	isInGame[key] = g

	dg, err := discordgo.New("Bot " + env["dgToken"])
	g.session = dg
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}
	dg.AddHandler(mcInGame)
	dg.AddHandler(rcInGame)
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}
	guild, _ := dg.Guild(m.GuildID)

	logmsg := time.Now().String() + ": " + guild.Name + " 에서 게임이 시작되었습니다."
	loggerLog.Println(logmsg)
	sendGuideMsg(dg, g)

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	dg.Close()
}

// 게임 시작 안내 메시지 전송.
func sendGuideMsg(s *discordgo.Session, g *game) {
	if s != nil {
		sendMsg, _ := s.ChannelMessageSendEmbed(g.chanID, embed.NewGenericEmbed("", ""))
		g.messageID = sendMsg.ID
	}
}

// 인게임 명령어 처리.
func mcInGame(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
}

// 인게임 버튼 상호작용 처리.
func rcInGame(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
	if r.UserID == s.State.User.ID {
		return
	}
}
