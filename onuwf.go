package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var (
	isUserIn map[string]bool
	isInGame map[string]*game
)

func init() {
	envInit()
	roleGuideInit()
	emojiInit()
	loggerInit()

	isUserIn = make(map[string]bool)
	isInGame = make(map[string]*game)
	/*
		conn, ctx := mongoConn()
		mongoDB := conn.Database("WF_Data")

		data := allData("people", mongoDB, ctx)
		fmt.Println(string(data))
	*/
}

func main() {
	dg, err := discordgo.New("Bot " + env["dgToken"])
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}
	dg.AddHandler(messageCreate)
	dg.AddHandler(messageReactionAdd)
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	dg.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.HasPrefix(m.Content, "ㅁ") {
		if m.Content == "ㅁ시작" {
			if isInGame[m.GuildID+m.ChannelID] != nil {
				s.ChannelMessageSend(m.ChannelID, "현재 게임이 진행중인 채널입니다.")
				return
			}
			if isUserIn[m.Author.ID] {
				s.ChannelMessageSend(m.ChannelID, "현재 게임을 플레이 중인 유저입니다.")
				return
			}
			go startGame(m)
		}
		if m.Content == "ㅁ강제종료" {
			if thisGame := isInGame[m.GuildID+m.ChannelID]; thisGame == nil {
				s.ChannelMessageSend(m.ChannelID, "현재 채널에서 진행중인 게임이 없습니다.")
				return
			} else {
				thisGame.session.Close()
				isInGame[m.GuildID+m.ChannelID] = nil
				s.ChannelMessageSend(m.ChannelID, "현재 채널에서 진행중인 게임을 강제종료합니다.")
				return
			}

		}
	}
}

func messageReactionAdd(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
	// 봇 자신이 선택한 이모지라면 무시.
	if r.UserID == s.State.User.ID {
		return
	}

	curChan, _ := s.Channel(r.ChannelID)
	// DM이 아닌 채널에서 사용한 이모지 중 게임이 시작되지 않은 길드에서 사용했다면 무시.
	if curChan.Type != discordgo.ChannelTypeDM && isInGame[r.GuildID+r.ChannelID] == nil {
		return
	}

	/*
		g := isInGame[r.GuildID+r.ChannelID]

		// 숫자 이모지 선택.
		for i := 1; i < 26; i++ {
			if r.Emoji.Name == emj[string(i)] {
				g.curState.pressNumBtn(s, r, i)
			}
		}

		// 쓰레기통 이모지 선택.
		if r.Emoji.Name == emj["DISCARD"] {
			g.curState.pressDisBtn(s, r)
		}

		// O 이모지 선택.
		if r.Emoji.Name == emj["YES"] {
			g.curState.pressYesBtn(s, r)
		}

		// X 이모지 선택.
		if r.Emoji.Name == emj["NO"] {
			g.curState.pressNoBtn(s, r)
		}

		// 왼쪽 화살표 선택.
		if r.Emoji.Name == emj["LEFT"] {
			g.curState.pressDirBtn(s, r, -1)
		}

		// 오른쪽 화살표 선택.
		if r.Emoji.Name == emj["RIGHT"] {
			g.curState.pressDirBtn(s, r, 1)
		}
	*/
}
