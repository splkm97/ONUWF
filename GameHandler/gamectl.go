package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	embed "github.com/clinet/discordgo-embed"
)

var (
	gid      *string
	curGame  *game
	isUserIn map[string]bool
)

func init() {
	envInit()
	roleGuideInit()
	emojiInit()
	isUserIn = make(map[string]bool)
	gid = flag.String("gid", "NO_GID", "실행한 길드의 고유값")
	cid := flag.String("cid", "NO_CID", "실행한 채널의 고유값")
	uid := flag.String("mid", "NO_MID", "게임을 시작한 방장의 유저 고유값")
	flag.Parse()
	curGame = newGame(*gid, *cid, *uid)
	isUserIn[*uid] = true
}

// GameHandler 는 새로운 게임이 시작될 때,
// 해당 게임을 관리하는 새 세션을 만든다.
// 이것은 강제종료 메시지를 수신하였을 때,
// 더 깔끔하게 수행할 수 있도록 한다.
func main() {
	dg, err := discordgo.New("Bot " + env["dgToken"])
	if err != nil {
		return
	}
	dg.AddHandler(messageCreate)
	dg.AddHandler(messageReactionAdd)
	err = dg.Open()
	if err != nil {
		return
	}
	sendGuideMsg(dg, curGame)

	sc := make(chan os.Signal, 1)
	curGame.killChan = sc
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	dg.Close()
}

// sendGuideMsg 함수는 게임 시작 안내 메시지를 전송한다.
func sendGuideMsg(s *discordgo.Session, g *game) {
	if s != nil {
		roleEmbed := embed.NewGenericEmbed("직업 추가", "1. 늑대인간 ...")
		roleMsg, _ := s.ChannelMessageSendEmbed(g.chanID, roleEmbed)
		g.roleAddMsgID = roleMsg.ID
		addRoleAddEmoji(s, roleMsg)
		enterEmbed := embed.NewGenericEmbed("게임 참가", "현재 참가 인원:\n")
		enterEmbed.Description += curGame.userList[0].nick
		enterEmbed.Footer = &discordgo.MessageEmbedFooter{Text: "⭕: 입장\n❌: 퇴장"}
		enterMsg, _ := s.ChannelMessageSendEmbed(g.chanID, enterEmbed)
		g.enterGameMsgID = enterMsg.ID
		addEnterGameEmoji(s, enterMsg)
	}
}

// addRoleAddEmoji 는 직업 추가 메시지에 안내 버튼을 연결해주는 함수입니다.
func addRoleAddEmoji(s *discordgo.Session, msg *discordgo.Message) {
	s.MessageReactionAdd(msg.ChannelID, msg.ID, emj["YES"])
	s.MessageReactionAdd(msg.ChannelID, msg.ID, emj["NO"])
	s.MessageReactionAdd(msg.ChannelID, msg.ID, emj["LEFT"])
	s.MessageReactionAdd(msg.ChannelID, msg.ID, emj["RIGHT"])
}

// addEnterGameEmoji 는 게임 입장 메시지에 안내 버튼을 연결해주는 함수입니다.
func addEnterGameEmoji(s *discordgo.Session, msg *discordgo.Message) {
	s.MessageReactionAdd(msg.ChannelID, msg.ID, emj["YES"])
	s.MessageReactionAdd(msg.ChannelID, msg.ID, emj["NO"])
}

// mcInGame 함수는 인게임 명령어 처리하는 이벤트 핸들러 함수입니다.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	if m.GuildID == curGame.guildID && m.ChannelID == curGame.chanID && isUserIn[m.Author.ID] {
		if m.Content == "ㅁ강제종료" {
			s.ChannelMessageSend(m.ChannelID, "게임을 강제종료합니다.")
			curGame.killChan <- os.Interrupt
		}
	}
}

// rcInGame 함수는 인게임 버튼 이모지 상호작용 처리를 위한 이벤트 핸들러 함수입니다.
func messageReactionAdd(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
	if r.UserID == s.State.User.ID {
		return
	}

	// DM이 아닌 곳에서 온 리액션 중,
	// 참가한 게임이 진행중인 길드, 채널 밖에서 달린 리액션을 무시한다.
	if curGame.guildID != r.GuildID && curGame.chanID != r.ChannelID {
		if ch, _ := s.Channel(r.ChannelID); ch.Type != discordgo.ChannelTypeDM {
			return
		}
	}

	g := curGame
	// 숫자 이모지 선택.
	for i := 1; i < 10; i++ {
		var ch rune
		ch = '0' + rune(i)
		emjID := "n" + string(ch)
		if r.Emoji.Name == emj[emjID] {
			g.curState.pressNumBtn(s, r, i)
		}
	}
	switch r.Emoji.Name {
	case emj["DISCARD"]:
		// 쓰레기통 이모지 선택.
		go g.curState.pressDisBtn(s, r)
	case emj["YES"]:
		// O 이모지 선택.
		go g.curState.pressYesBtn(s, r)
	case emj["NO"]:
		// X 이모지 선택.
		go g.curState.pressNoBtn(s, r)
	case emj["LEFT"]:
		// 왼쪽 화살표 선택.
		go g.curState.pressDirBtn(s, r, -1)
	case emj["RIGHT"]:
		// 오른쪽 화살표 선택.
		go g.curState.pressDirBtn(s, r, 1)
	}
	if r.GuildID == curGame.guildID && r.ChannelID == curGame.chanID && (r.MessageID == curGame.enterGameMsgID || r.MessageID == curGame.roleAddMsgID) {
		if isUserIn[r.UserID] {
			s.MessageReactionRemove(r.ChannelID, r.MessageID, r.Emoji.Name, r.UserID)
		}
	}

}
