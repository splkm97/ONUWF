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
	Action(tar targetObject, player *user, g *game)
	String() string
}

type roleFactory struct {
}

func (rf *roleFactory) generateRole(num int) (r role) {
	switch num {
	case 1:
		r = &roleSentinel{}
		/*
			case 2:
				r = roleDoppelganger{}
			case 3:
				r = roleWerewolf{}
			case 4:
				r = roleAlphawolf{}
			case 5:
				r = roleMisticwolf{}
			case 6:
				r = roleMinion{}
			case 7:
				r = roleFreemasonry{}
			case 8:
				r = roleSeer{}
			case 9:
				r = roleApprenticeseer{}
			case 10:
				r = roleParanormalinvestigator{}
			case 11:
				r = roleRober{}
			case 12:
				r = roleWitch{}
			case 13:
				r = roleTroublemaker{}
			case 14:
				r = roleVillageidiot{}
			case 15:
				r = roleDrunk{}
			case 16:
				r = roleInsomniac{}
			case 17:
				r = roleRevealer{}
			case 18:
				r = roleTanner{}
			case 19:
				r = roleHunter{}
			case 20:
				r = roleBodygaurd{}
			case 21:
				r = roleVillager{}
			case 22:
				r = roleDreamwolf{}
		*/

	}
	return r
}

// targetObject 는 각 직업의 특수 능력이
// 적용되어야 하는 대상을 구분하는 기준으로 사용된다.
type targetObject struct {
	actionType int
	uid1       string
	uid2       string
	disRoleIdx int
}

type roleSentinel struct {
	role
}

func (r roleSentinel) Action(tar targetObject, player *user, g *game) {

}

func (r roleSentinel) String() string {
	return "수호자"
}
