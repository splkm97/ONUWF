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
