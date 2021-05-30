package main

type state interface {
	// 사용자 인원수 3 ~ 26
	// num: 0 ~ 23
	// pressNumBtn 사용자가 숫자 이모티콘을 눌렀을 때 state에서 하는 동작
	pressNumBtn func(s *discordgo.Session, r *discordgo.MessageReactionAdd, num int)

	// pressDisBtn 사용자가 버려진 카드 이모티콘을 눌렀을 때 state에서 하는 동작
	pressDisBtn func(s *discordgo.Session, r *discordgo.MessageReactionAdd)

	// pressYesBtn 사용자가 yes 이모티콘을 눌렀을 때 state에서 하는 동작
	pressYesBtn func(s *discordgo.Session, r *discordgo.MessageReactionAdd)

	// pressNoBtn 사용자가 No 이모티콘을 눌렀을 때 state에서 하는 동작
	pressNoBtn func(s *discordgo.Session, r *discordgo.MessageReactionAdd)

	// 좌 -1, 우 1
	// pressDirBtn 사용자가 좌우 방향 이모티콘을 눌렀을 때 state에서 하는 동작
	pressDirBtn func(s *discordgo.Session, r *discordgo.MessageReactionAdd, dir int)

	// sendFinish 사용자가 종료 메세지를 보냈을 때 state에서 하는 동작
	sendFinish func(s *discordgo.Session, m *discordgo.MessageCreate)
}

type StatePrepare struct {
	// state에서 가지고 있는 game
	g	*game
}

// appendRole 현재 게임에 직업을 추가
func (sPrepare StatePrepare) appendRole(item role) {
	sPrepare.g.roleSeq = append(sPrepare.g.roleSeq, item)
}

// removeRole 현재 게임에 직업을 삭제
func (sPrepare StatePrepare) removeRole(item role) {
	if index := findRoleIdx(item, sPrepare.g.roleSeq); index != -1 {
		sPrepare.g.roleSeq = append(sPrepare.g.roleSeq[:index], sPrepare.g.roleSeq[index+1:]...)
	}
}

// 사용자가 숫자 이모티콘을 눌렀을 때 StatePrepare에서 하는 동작
func (sPrepare StatePrepare) pressNumBtn(s *discordgo.Session, r *discordgo.MessageReactionAdd, num int) {
	var var_role role
	if num == 1 {
		var_role = roleSentinel{}
	} else if num == 2 {
		var_role = roleDoppelganger{}
	} else if num == 3 {
		var_role = roleWerewolf}
	} else if num == 4 {
		var_role = roleAlphawolf{}
	} else if num == 5 {
		var_role = roleMisticwolf{}
	} else if num == 6 {
		var_role = roleMinion{}
	} else if num == 7 {
		var_role = roleFreemasonry{}
	} else if num == 8 {
		var_role = roleSeer{}
	} else if num == 9 {
		var_role = roleApprenticeseer{}
	} else if num == 10 {
		var_role = roleParanormalinvestigator{}
	} else if num == 11 {
		var_role = roleRober{}
	} else if num == 12 {
		var_role = roleWitch{}
	} else if num == 13 {
		var_role = roleTroublemaker{}
	} else if num == 14 {
		var_role = roleVillageidiot{}
	} else if num == 15 {
		var_role = roleDrunk{}
	} else if num == 16 {
		var_role = roleInsomniac{}
	} else if num == 17 {
		var_role = roleRevealer{}
	} else if num == 18 {
		var_role = roleTanner{}
	} else if num == 19 {
		var_role = roleHunter{}
	} else if num == 20 {
		var_role = roleBodugaurd{}
	} else if num == 21 {
		var_role = roleVillager{}
	} else if num == 22 {
		var_role = roleDreamwolf{}
	}
	sPrepare.appendRole(role)
	if len(g.roleView) == len(g.userList)+3 {
		g.state = StatePlayable{g: g}
	}
}

// 사용자가 버려진 카드 이모티콘을 눌렀을 때 StatePrepare에서 하는 동작
func (sPrepare StatePrepare) pressDisBtn(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
	//do nothing
}

// 사용자가 yes 이모티콘을 눌렀을 때 StatePrepare에서 하는 동작
func (sPrepare StatePrepare) pressYesBtn(s *discordgo.Session, r *discordgo.MessageReactionAdd) {

}

// 사용자가 No 이모티콘을 눌렀을 때 StatePrepare에서 하는 동작
func (sPrepare StatePrepare) pressNoBtn(s *discordgo.Session, r *discordgo.MessageReactionAdd) {

}

// 좌 -1, 우 1
// 사용자가 좌우 방향 이모티콘을 눌렀을 때 StatePrepare에서 하는 동작
func (sPrepare StatePrepare) pressDirBtn(s *discordgo.Session, r *discordgo.MessageReactionAdd, dir int) {

}

// 사용자가 종료 메세지를 보냈을 때 StatePrepare에서 하는 동작
func (sPrepare StatePrepare) sendFinish(s *discordgo.Session, m *discordgo.MessageCreate) {


}