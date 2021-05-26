package main

// 게임 진행을 위한 정보를 담고 있는 스트럭처
type game struct {
	// 현재 게임이 진행중인 서버의 GID
	guildID string

	// 현재 게임이 진행중인 채널의 CID
	chanID string

	// 현재 게임에서 설정된 직업들의 목록
	roleIDlist []int

	// 현재 게임의 참가자 UID
	userIDlist []string

	// 현재 게임의 진행시점
	curState state

	// Role을 User별로 매핑시킨 인덱스 테이블
	roleIdxTable [][]bool
	oriRoleIdxTable [][]bool

	// 게임 진행 상황을 기록하는 로그 메시지 배열
	logMsg []string
}

