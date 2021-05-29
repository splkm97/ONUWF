package main

/*
type state struct {
	// state에서 가지고 있는 game
	g *game

	// 사용자 인원수 3 ~ 26
	// num: 0 ~ 23
	// 사용자가 숫자 이모티콘을 눌렀을 때 state에서 하는 동작
	pressNumBtn func(s *discordgo.Session, r *discordgo.MessageReactionAdd, num int)

	// 사용자가 버려진 카드 이모티콘을 눌렀을 때 state에서 하는 동작
	pressDisBtn func(s *discordgo.Session, r *discordgo.MessageReactionAdd)

	// 사용자가 yes 이모티콘을 눌렀을 때 state에서 하는 동작
	pressYesBtn func(s *discordgo.Session, r *discordgo.MessageReactionAdd)

	// 사용자가 No 이모티콘을 눌렀을 때 state에서 하는 동작
	pressNoBtn func(s *discordgo.Session, r *discordgo.MessageReactionAdd)

	// 좌 -1, 우 1
	// 사용자가 좌우 방향 이모티콘을 눌렀을 때 state에서 하는 동작
	pressDirBtn func(s *discordgo.Session, r *discordgo.MessageReactionAdd, dir int)

	// 사용자가 종료 메세지를 보냈을 때 state에서 하는 동작
	sendFinish func(s *discordgo.Session, m *discordgo.MessageCreate)
}

// StatePrepare
type StatePrepare struct {
	// state를 상속함
	state
}

// 사용자가 숫자 이모티콘을 눌렀을 때 StatePrepare에서 하는 동작
func (sPrepare *StatePrepare) pressNumBtn(s *discordgo.Session, r *discordgo.MessageReactionAdd, num int) {
	sPrepare.g.addRole(werewolf{})
	//직업추가
	if len(g.roleView) == len(g.userList)+3 {
		g.state = StatePlayable{g: g}
	}
}

// 사용자가 버려진 카드 이모티콘을 눌렀을 때 StatePrepare에서 하는 동작
func (sPrepare *StatePrepare) pressDisBtn(s *discordgo.Session, r *discordgo.MessageReactionAdd) {

}

// 사용자가 yes 이모티콘을 눌렀을 때 StatePrepare에서 하는 동작
func (sPrepare *StatePrepare) pressYesBtn(s *discordgo.Session, r *discordgo.MessageReactionAdd) {

}

// 사용자가 No 이모티콘을 눌렀을 때 StatePrepare에서 하는 동작
func (sPrepare *StatePrepare) pressNoBtn(s *discordgo.Session, r *discordgo.MessageReactionAdd) {

}

// 좌 -1, 우 1
// 사용자가 좌우 방향 이모티콘을 눌렀을 때 StatePrepare에서 하는 동작
func (sPrepare *StatePrepare) pressDirBtn(s *discordgo.Session, r *discordgo.MessageReactionAdd, dir int) {

}

// 사용자가 종료 메세지를 보냈을 때 StatePrepare에서 하는 동작
func (sPrepare *StatePrepare) sendFinish(s *discordgo.Session, m *discordgo.MessageCreate) {


}
*/
