package main

import (
	"github.com/bwmarrin/discordgo"
	embed "github.com/clinet/discordgo-embed"
)

type StatePlayable struct {
	// state에서 가지고 있는 game
	g *game
}

// pressNumBtn 사용자가 숫자 이모티콘을 눌렀을 때 StatePlayable에서 하는 동작
func (sPrepare *StatePlayable) pressNumBtn(s *discordgo.Session, r *discordgo.MessageReactionAdd, num int) {

}

// pressDisBtn 사용자가 버려진 카드 이모티콘을 눌렀을 때 StatePlayable에서 하는 동작
func (sPrepare *StatePlayable) pressDisBtn(s *discordgo.Session, r *discordgo.MessageReactionAdd) {

}

// pressYesBtn 사용자가 yes 이모티콘을 눌렀을 때 StatePlayable에서 하는 동작
func (sPrepare *StatePlayable) pressYesBtn(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
	// 입장 메세지에서 리액션한거라면
	if r.MessageID == sPrepare.g.enterGameMsgID {
		s.ChannelMessageSendEmbed(sPrepare.g.chanID, embed.NewGenericEmbed("시작가능", ""))
		// 직업추가 메세지에서 리액션한거라면
	} else if r.MessageID == sPrepare.g.roleAddMsgID {

	}
}

// pressNoBtn 사용자가 No 이모티콘을 눌렀을 때 StatePlayable에서 하는 동작
func (sPrepare *StatePlayable) pressNoBtn(s *discordgo.Session, r *discordgo.MessageReactionAdd) {

}

// pressDirBtn 좌 -1, 우 1 사용자가 좌우 방향 이모티콘을 눌렀을 때 StatePlayable에서 하는 동작
func (sPrepare *StatePlayable) pressDirBtn(s *discordgo.Session, r *discordgo.MessageReactionAdd, dir int) {

}

// sendFinish 사용자가 종료 메세지를 보냈을 때 StatePlayable에서 하는 동작
func (sPrepare *StatePlayable) sendFinish(s *discordgo.Session, m *discordgo.MessageCreate) {

}
