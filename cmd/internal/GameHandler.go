// +build linux,amd64,go1.15,!cgo

package main

import (
	"flag"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	embed "github.com/clinet/discordgo-embed"
)

var (
	gid        *string
	curSession *discordgo.Session
	curGame    *gamedata.Game
	isUserIn   map[string]bool
	rf         RoleFactory
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
	rf = roleFactory{}
	dg.AddHandler(messageCreate)
	dg.AddHandler(messageReactionAdd)
	err = dg.Open()
	if err != nil {
		return
	}
	curGame.setUserByID(dg, curGame.masterID)
	curSession = dg
	sendGuideMsg(dg)

	sc := make(chan os.Signal, 1)
	curGame.killChan = sc
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	dg.Close()
}

func roleCount(roleToFind Role, roleView []Role) int {
	cnt := 0
	findRoleName := roleToFind.String()
	for _, tmpRole := range roleView {
		if findRoleName == tmpRole.String() {
			cnt++
		}
	}
	return cnt
}

// newRoleEmbed 함수는 role guide와 현재 게임에 추가된 직업 / 게임의 참여중인 인원수 + 3 임베드를 만든다
func newRoleEmbed(rgIndex int, g *Game) *embed.Embed {
	roleEmbed := embed.NewEmbed()
	roleEmbed.SetTitle("직업 추가")
	roleEmbed.AddField(rg[rgIndex].RoleName, strings.Join(rg[rgIndex].RoleGuide, "\n"))
	roleStr := ""
	if len(g.roleView) == 0 {
		roleStr += "*추가된 직업이 없습니다.*"
	}
	for _, item := range g.roleView {
		cnt := roleCount(item, g.roleView)
		roleStr += item.String() + " " + strconv.Itoa(cnt) + "개"
		if cnt == rg[rgIndex].Max {
			roleStr += " 최대"
		}
		roleStr += "\n"
	}
	roleEmbed.AddField("추가된 직엄", roleStr)
	roleEmbed.SetFooter("현재 인원에 맞는 직업 수: " + strconv.Itoa(len(g.roleView)) + " / " + strconv.Itoa(len(g.userList)+3))
	return roleEmbed
}

// newEnterEmbed 함수는 게임 참여자 목록 임베드를 만든다
func newEnterEmbed(g *Game) *embed.Embed {
	enterEmbed := embed.NewEmbed()
	enterEmbed.SetTitle("게임 참가")
	enterEmbed.AddField("", "현재 참가 인원: "+strconv.Itoa(len(g.userList))+"명\n")
	enterStr := ""
	if len(g.userList) == 0 {
		enterStr += "*참가자가 없습니다.*"
	}
	for _, item := range g.userList {
		enterStr += "`" + item.nick + "`\n"
	}
	enterEmbed.AddField("참가자 목록", enterStr)
	enterEmbed.SetFooter("⭕: 입장\n❌: 퇴장")
	return enterEmbed
}

// sendGuideMsg 함수는 게임 시작 안내 메시지를 전송한다.
func sendGuideMsg(s *discordgo.Session) {
	if s != nil {
		roleEmbed := newRoleEmbed(0, curGame)
		roleMsg, _ := s.ChannelMessageSendEmbed(curGame.chanID, roleEmbed.MessageEmbed)
		curGame.roleAddMsgID = roleMsg.ID
		addRoleAddEmoji(s, roleMsg)
		enterEmbed := newEnterEmbed(curGame)
		enterMsg, _ := s.ChannelMessageSendEmbed(curGame.chanID, enterEmbed.MessageEmbed)
		curGame.enterGameMsgID = enterMsg.ID
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
		if m.Content == "ㅁ테스트" {
			curGame.sendVoteMsg(s)
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
		s.MessageReactionRemove(r.ChannelID, r.MessageID, r.Emoji.Name, r.UserID)
	}

}
