package main

import "github.com/bwmarrin/discordgo"

type state interface {
	// 사용자 인원수 3 ~ 26
	// num: 0 ~ 23
	// pressNumBtn 사용자가 숫자 이모티콘을 눌렀을 때 state에서 하는 동작
	pressNumBtn(s *discordgo.Session, r *discordgo.MessageReactionAdd, num int)

	// pressDisBtn 사용자가 버려진 카드 이모티콘을 눌렀을 때 state에서 하는 동작
	pressDisBtn(s *discordgo.Session, r *discordgo.MessageReactionAdd)

	// pressYesBtn 사용자가 yes 이모티콘을 눌렀을 때 state에서 하는 동작
	pressYesBtn(s *discordgo.Session, r *discordgo.MessageReactionAdd)

	// pressNoBtn 사용자가 No 이모티콘을 눌렀을 때 state에서 하는 동작
	pressNoBtn(s *discordgo.Session, r *discordgo.MessageReactionAdd)

	// pressDirBtn 좌 -1, 우 1 사용자가 좌우 방향 이모티콘을 눌렀을 때 state에서 하는 동작
	pressDirBtn(s *discordgo.Session, r *discordgo.MessageReactionAdd, dir int)

	// sendFinish 사용자가 종료 메세지를 보냈을 때 state에서 하는 동작
	sendFinish(s *discordgo.Session, m *discordgo.MessageCreate)
}

type StatePrepare struct {
	// state에서 가지고 있는 game
	g *game

	// factory 에서 쓰이게 될 role index
	roleIndex	int

	// roleFactory
	rf	roleFactory
}

// removeRole 현재 게임에 직업을 삭제
func (sPrepare StatePrepare) removeRole(item role) {
	if index := findRoleIdx(item, sPrepare.g.roleSeq); index != -1 {
		sPrepare.g.roleSeq = append(sPrepare.g.roleSeq[:index], sPrepare.g.roleSeq[index+1:]...)
	}
}

// pressNumBtn 사용자가 숫자 이모티콘을 눌렀을 때 StatePrepare에서 하는 동작
func (sPrepare StatePrepare) pressNumBtn(s *discordgo.Session, r *discordgo.MessageReactionAdd, num int) {
	// do nothing
}

// pressDisBtn 사용자가 버려진 카드 이모티콘을 눌렀을 때 StatePrepare에서 하는 동작
func (sPrepare StatePrepare) pressDisBtn(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
	// do nothing
}

// pressYesBtn 사용자가 yes 이모티콘을 눌렀을 때 StatePrepare에서 하는 동작
func (sPrepare StatePrepare) pressYesBtn(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
	if r.MessageID == sPrepare.g.enterGameMsgID {
		user := {userID: r.UserID, nick: s.User(r.UserID).Username, chanID: r.ChannelID, dmChanID: s.UserChannelCreate(r.UserID).ID}
		userList = append(userList, user)
	} else if r.MessageID == sPrepare.g.roleAddMsgID {
		// roleFactory에서 현재 roleindex 위치 값을 받아
		sPrepare.g.roleSeq = append(sPrepare.g.roleSeq, rf.make(sPrepare.rolIndex))
		if len(g.roleView) == len(g.userList)+3 {
			g.state = StatePlayable{g: g}
		}
	}
}

// pressNoBtn 사용자가 No 이모티콘을 눌렀을 때 StatePrepare에서 하는 동작
func (sPrepare StatePrepare) pressNoBtn(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
	sPrepare.removeRole(sPrepare.rolIndex)
}

// pressDirBtn 좌 -1, 우 1 사용자가 좌우 방향 이모티콘을 눌렀을 때 StatePrepare에서 하는 동작
func (sPrepare StatePrepare) pressDirBtn(s *discordgo.Session, r *discordgo.MessageReactionAdd, dir int) {
	// 게임 시작

	// roleindex 증감
}

// sendFinish 사용자가 종료 메세지를 보냈을 때 StatePrepare에서 하는 동작
func (sPrepare StatePrepare) sendFinish(s *discordgo.Session, m *discordgo.MessageCreate) {

}
