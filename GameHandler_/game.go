package main

import (
	"os"

	"github.com/bwmarrin/discordgo"
)

// game 구조체는 게임 진행을 위한 정보를 담고 있는 스트럭처
type game struct {
	// 현재 게임이 진행중인 서버의 GID
	guildID string

	// 현재 게임이 진행중인 채널의 CID
	chanID string

	// 현재 게임의 세션 주소
	session *discordgo.Session

	roleAddMsgID   string
	enterGameMsgID string

	// 게임을 생성한 방장의 UID
	masterID string

	// 현재 게임의 참가자들
	userList []user

	// 현재 게임에서 순서대로 추가, 중복제거 된 직업들의 목록
	roleSeq []role

	// 현재 게임에서 사용중인 사용자에게 보여줄 중복 정렬된 직업들의 목록
	roleView []role

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

	killChan chan os.Signal
}

func newGame(gid, cid, mid string) (g *game) {
	g = &game{}
	g.guildID = gid
	g.chanID = cid
	g.masterID = mid
	g.userList = make([]user, 0)
	g.roleSeq = make([]role, 0)
	g.disRole = make([]role, 0)
	g.curState = &StatePrepare{g, 1, nil, nil}
	g.logMsg = make([]string, 0)
	return
}

func (g *game) setUserByID(s *discordgo.Session, uid string) {
	var newone user
	newone.userID = uid
	dgUser, _ := s.User(uid)
	newone.nick = dgUser.Username
	newone.chanID = g.chanID
	uChan, _ := s.UserChannelCreate(uid)
	newone.dmChanID = uChan.ID
	g.userList = append(g.userList, newone)
}

// UID 로 user 인스턴스를 구하는 함수
func (g *game) findUserByUID(uid string) (target *user) {
	for i, item := range g.userList {
		if item.userID == uid {
			return &g.userList[i]
		}
	}
	return nil
}

// 게임 로그에 메시지를 쌓는 함수.
func (g *game) appendLog(msg string) {
	if g.logMsg == nil {
		g.logMsg = make([]string, 0)
	}
	g.logMsg = append(g.logMsg, msg)
}

// 유저의 직업을 반환
func (g *game) getRole(uid string) role {
	loop := len(g.roleSeq)
	idx := findUserIdx(uid, g.userList)

	for i := 0; i < loop; i++ {
		if g.roleIdxTable[idx][i] {
			return g.roleSeq[i]
		}
	}
	return nil
}

// 유저의 직업을 업데이트
func (g *game) setRole(uid string, item role) {
	userIdx := findUserIdx(uid, g.userList)
	roleIdx := findRoleIdx(item, g.roleSeq)
	loop := len(g.roleSeq)

	for i := 0; i < loop; i++ {
		g.roleIdxTable[userIdx][i] = false
	}
	g.roleIdxTable[userIdx][roleIdx] = true
}

// 버려진 직업을 업데이트
func (g *game) setDisRole(disRoleIdx int, item role) {
	g.disRole[disRoleIdx] = item
}

// 두 유저의 직업을 서로 교환
func (g *game) swapRoleFromUser(uid1, uid2 string) {
	role1 := g.getRole(uid1)
	role2 := g.getRole(uid2)
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
func (g *game) getRoleUsers(find role) (users []user) {
	result := make([]user, 0)
	loop := len(g.userList)

	idx := findRoleIdx(find, g.roleSeq)

	for i := 0; i < loop; i++ {
		if g.roleIdxTable[i][idx] {
			result = append(result, g.userList[i])
		}
	}

	return result
}

// 모든 사람들의 직업을 입장순서별로 한칸 회전.
func (g *game) rotateAllUserRole() {
	loop := len(g.userList)

	tmpRole := g.getRole(g.userList[loop-1].userID)
	for i := loop - 1; i > 0; i++ {
		item := g.getRole(g.userList[i-1].userID)
		g.setRole(g.userList[i].userID, item)
	}
	g.setRole(g.userList[0].userID, tmpRole)
}

// 유저에게 특수권한 부여
func (g *game) setPower(power int) {
	// TODO 내부 구현.
}

// 특정 유저의 직업을 복사.
func (g *game) copyRole(destUID, srcUID string) {
	srcRole := g.getRole(srcUID)
	g.setRole(destUID, srcRole)
}

// 유저의 인덱스 찾기를 위한 함수
func findUserIdx(uid string, target []user) int {
	for i, item := range target {
		if uid == item.userID {
			return i
		}
	}
	return -1
}

// 직업의 인덱스 찾기를 위한 함수
func findRoleIdx(r role, target []role) int {
	for i, item := range target {
		if r.String() == item.String() {
			return i
		}
	}
	return -1
}
