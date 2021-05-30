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

// startGame 은 새로운 게임이 시작될 때,
// 해당 게임을 관리하는 새 세션을 만든다.
// 이것은 강제종료 메시지를 수신하였을 때,
// 더 깔끔하게 수행할 수 있도록 한다.
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
	g.killChan = sc
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	dg.Close()
}

// sendGuideMsg 함수는 게임 시작 안내 메시지를 전송한다.
func sendGuideMsg(s *discordgo.Session, g *game) {
	if s != nil {
		roleMsg, _ := s.ChannelMessageSendEmbed(g.chanID, embed.NewGenericEmbed("직업 추가", "1. 늑대인간 ..."))
		g.roleAddMsgID = roleMsg.ID
		addRoleAddEmoji(s, roleMsg)
		enterMsg, _ := s.ChannelMessageSendEmbed(g.chanID, embed.NewGenericEmbed("게임 참가", ": 입장\n: 퇴장"))
		g.enterGameMsgID = enterMsg.ID
		addEnterGameEmoji(s, enterMsg)
	}
}

func addRoleAddEmoji(s *discordgo.Session, msg *discordgo.Message) {
	s.MessageReactionAdd(msg.ChannelID, msg.ID, emj["LEFT"])
	s.MessageReactionAdd(msg.ChannelID, msg.ID, emj["YES"])
	s.MessageReactionAdd(msg.ChannelID, msg.ID, emj["NO"])
	s.MessageReactionAdd(msg.ChannelID, msg.ID, emj["RIGHT"])
}

func addEnterGameEmoji(s *discordgo.Session, msg *discordgo.Message) {
	s.MessageReactionAdd(msg.ChannelID, msg.ID, emj["YES"])
	s.MessageReactionAdd(msg.ChannelID, msg.ID, emj["NO"])
}

// mcInGame 함수는 인게임 명령어 처리하는 이벤트 핸들러 함수이다.
func mcInGame(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	if m.Content == "ㅇㅋ" {
		for {
			s.ChannelMessageSend(m.ChannelID, "ㅇㅋㅇㅋ")
			time.Sleep(time.Second * 3)
		}
	}
	if m.Content == "ㄴㄴ" {
		s.ChannelMessageSend(m.ChannelID, "ㅇㅋㅇㅋ")
	}
}

// rcInGame 함수는 인게임 버튼 이모지 상호작용 처리를 위한 이벤트 핸들러 함수이다.
func rcInGame(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
	if r.UserID == s.State.User.ID {
		return
	}

	// 게임이 시작되지 않았으면 리액션을 무시한다.
	if isInGame[r.GuildID+r.ChannelID] == nil {
		return
	}

	g := isInGame[r.GuildID+r.ChannelID]
	// 숫자 이모지 선택.
	for i := 1; i < 10; i++ {
		var ch rune
		ch = '0' + rune(i)
		emjID := "n" + string(ch)
		if r.Emoji.Name == emj[emjID] {
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
}
