package main

// 각 사용자들의 정보를 담고 있는 스트럭처
type user struct {
	// 각 유저의 UID
	userID string

	// 각 유저의 닉네임
	nick string

	// 각 유저가 속한 게임이 진행중인 채널 ID
	chanID string

	// 각 유저의 DM 채널 ID
	dmChanID string
}

// 각 직업들의 정보를 담고 있는 스트럭처
type role interface {
	// 각 직업별 행동 함수를 다르게 정의하기 위한 함수 선언
	Action(uid1, uid2 string, disRoleIdx int, player *user, g *game)
	String() string
}

type roleFactory struct {

}

func (rf roleFactory) make(i int) r role {
	if num == 1 {
		r = roleSentinel{}
	} else if num == 2 {
		r = roleDoppelganger{}
	} else if num == 3 {
		r = roleWerewolf}
	} else if num == 4 {
		r = roleAlphawolf{}
	} else if num == 5 {
		r = roleMisticwolf{}
	} else if num == 6 {
		r = roleMinion{}
	} else if num == 7 {
		r = roleFreemasonry{}
	} else if num == 8 {
		r = roleSeer{}
	} else if num == 9 {
		r = roleApprenticeseer{}
	} else if num == 10 {
		r = roleParanormalinvestigator{}
	} else if num == 11 {
		r = roleRober{}
	} else if num == 12 {
		r = roleWitch{}
	} else if num == 13 {
		r = roleTroublemaker{}
	} else if num == 14 {
		r = roleVillageidiot{}
	} else if num == 15 {
		r = roleDrunk{}
	} else if num == 16 {
		r = roleInsomniac{}
	} else if num == 17 {
		r = roleRevealer{}
	} else if num == 18 {
		r = roleTanner{}
	} else if num == 19 {
		r = roleHunter{}
	} else if num == 20 {
		r = roleBodugaurd{}
	} else if num == 21 {
		r = roleVillager{}
	} else if num == 22 {
		r = roleDreamwolf{}
	}
	return r
}
