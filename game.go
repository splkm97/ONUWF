package main

// 게임 진행을 위한 정보를 담고 있는 스트럭처
type game struct {
	// 현재 게임이 진행중인 서버의 GID
	guildID string

	// 현재 게임이 진행중인 채널의 CID
	chanID string

	// 게임을 생성한 방장의 UID
	masterID string

	// 현재 게임의 참가자들
	userlist []user

	// 현재 게임에서 설정된 직업들의 목록
	rolelist []role

	// 현재 게임의 진행시점
	curState state

	// Role을 User별로 매핑시킨 인덱스 테이블
	// <usage : roleIdxTable[userIdx][roleIdx]>
	roleIdxTable    [][]bool
	oriRoleIdxTable [][]bool

	// 게임에서 버려진 직업 목록
	disRole []role

	// 게임 진행 상황을 기록하는 로그 메시지 배열
	logMsg []string
}

func newGame(gid, cid, mid string) (g *game) {
	g = &game{}
	g.guildID = gid
	g.chanID = cid
	g.masterID = mid
	g.userlist = make([]user, 0)
	g.rolelist = make([]role, 0)
	g.disRole = make([]role, 0)
	g.curState = StatePrepare{}
	g.logMsg = make([]string, 0)
}

// UID 로 user 인스턴스를 구하는 함수
func (g *game) findUserByUID(uid string) (target *user) {
	for i, item := range g.userlist {
		if item.userID == uid {
			return &g.userlist[i]
		}
	}
	return nil
}

// 게임 로그에 메시지를 쌓는 함수.
func (g *game) appendLog(msg string) {
	if logMsg == nil {
		logMsg = make([]string, 0)
	}
	logMsg = append(logMsg, msg)
}

// 유저의 직업을 반환
func (g *game) getRole(uid string) role {
	loop := len(g.rolelist) - 3

	idx := findUserIdx(uid, g.userlist)

	for i := 0; i < loop; i++ {
		if g.roleIdxTable[idx][i] == true {
			return g.rolelist[i]
		}
	}

	return nil
}

// 유저의 직업을 업데이트
func (g *game) setRole(uid string, item role) {
	userIdx := findUserIdx(uid, g.userlist)
	roleIdx := findRoleIdx(item, g.rolelist)
	loop := len(g.rolelist)

	for i := 0; i < loop; i++ {
		g.roleIdxTable[userIdx][i] = false
	}
	g.roleIdxTable[userIdx][roleIdx] = true
}

// 버려진 직업을 업데이트
func (g *game) setDisRole(disRoleIdx int, item role) {
	g.disRole[disRoleIdx] = role
}

// 두 유저의 직업을 서로 교환
func (g *game) swapRoleFromUser(uid1, uid2 string) {
	role1 := g.getRole(uid1, g)
	role2 := g.getRole(uid2, g)
	g.setRole(uid1, role2)
	g.setRole(uid2, role1)
}

// 버려진 직업 중 하나 확인.
func (g *game) getDisRole(disRoleIdx int) role {
	return g.disRole[disRoleIdx]
}

// 유저 직업과 버려진 직업을 교환.
func (g *game) swapRoleFromDiscard(uid string, disRoleIdx int) {
	role1 := g.getDisRole(disRoleIdx)
	role2 := g.getRole(uid)
	g.setRole(uid, role1)
	g.setDisRole(disRoleIdx, role2)
}

// 특정 직업의 유저 목록 반환.
func (g *game) getRoleUser(find *role) (users []user) {
	result := make([]user, 0)
	loop := len(g.userlist)

	idx := findRoleIdx(*find, g.rolelist)

	for i := 0; i < loop; i++ {
		if g.roleIdxTable[i][idx] == true {
			result = append(result, g.userlist[i])
		}
	}

	return result
}

// 모든 사람들의 직업을 입장순서별로 한칸 회전.
func (g *game) rotateAllUserRole() {
	loop := len(g.userlist)

	tmpRole := g.getRole(g.userlist[loop-1].userID)
	for i := loop - 1; i > 0; i++ {
		item := g.getRole(g.userlist[i-1].userID)
		g.setRole(g.userlist[i].userID, item)
	}
	g.setRole(g.userlist[0].userID, tmpRole)
}

// 유저에게 특수권한 부여
func (g *game) setPower(power int) {
	// TODO 내부 구현.
}

// 특정 유저의 직업을 복사.
func (g *game) copyRole(destUID, srcUID string) {
	srcRole := g.getRole(srcUID, g)
	srcIdx := findUserIdx(srcUID, g.userlist)
	destIdx := findUserIdx(destUID, g.userlist)
	for i := 0; i < len(g.rolelist); i++ {
		if roleIdxTable[srcIdx][i] == true {
			roleIdxTable[destIdx][i] = true
		} else {
			roleIdxTable[destIdx][i] = false
		}
	}
}

// 유저의 인덱스 찾기를 위한 함수
func findUserIdx(uid string, target []user) int {
	for i, item := range target {
		if str == item.userID {
			return i
		}
	}
	return -1
}

// 직업의 인덱스 찾기를 위한 함수
func findRoleIdx(r role, target []role) int {
	for i, item := range target {
		if g.String() == item.String() {
			return i
		}
	}
	return -1
}
