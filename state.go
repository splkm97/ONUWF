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
	roleIndex int

	// roleFactory
	rf roleFactory
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
		if findUserIdx(r.UserID, sPrepare.g.userList) == -1 {
			//roleMsg, _ := s.ChannelMessageSendEmbed(g.chanID, embed.NewGenericEmbed("직업 추가", "1. 늑대인간 ..."))

		} else {
			// user 만들어서 userList에 append()
			userNick, _ := s.User(r.UserID)
			userDM, _ := s.UserChannelCreate(r.UserID)
			u := user{userID: r.UserID, nick: userNick.Username, chanID: r.ChannelID, dmChanID: userDM.ID}
			sPrepare.g.userList = append(sPrepare.g.userList, u)
		}
	} else if r.MessageID == sPrepare.g.roleAddMsgID {
		// roleFactory에서 현재 roleindex 위치 값을 받아 role 추가
		sPrepare.g.roleSeq = append(sPrepare.g.roleSeq, sPrepare.rf.generateRole(sPrepare.roleIndex))
	}
}

// pressNoBtn 사용자가 No 이모티콘을 눌렀을 때 StatePrepare에서 하는 동작
func (sPrepare StatePrepare) pressNoBtn(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
	if index := findRoleIdx(sPrepare.rf.generateRole(sPrepare.roleIndex), sPrepare.g.roleSeq); index != -1 {
		sPrepare.g.roleSeq = append(sPrepare.g.roleSeq[:index], sPrepare.g.roleSeq[index+1:]...)
	}
}

// pressDirBtn 좌 -1, 우 1 사용자가 좌우 방향 이모티콘을 눌렀을 때 StatePrepare에서 하는 동작
func (sPrepare StatePrepare) pressDirBtn(s *discordgo.Session, r *discordgo.MessageReactionAdd, dir int) {
	if r.MessageID == sPrepare.g.enterGameMsgID {
		// 게임 시작
		if dir == 1 {
			if len(sPrepare.g.roleSeq) == len(sPrepare.g.userList) + 3 {
				g.state = StatePlayable{g: sPrepare.g}
			} else {

			}
		}
	} else if r.MessageID == sPrepare.g.roleAddMsgID {
		// roleindex 증감
		
	}
}

// sendFinish 사용자가 종료 메세지를 보냈을 때 StatePrepare에서 하는 동작
func (sPrepare StatePrepare) sendFinish(s *discordgo.Session, m *discordgo.MessageCreate) {

}
