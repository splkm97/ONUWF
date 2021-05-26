package main

// 각 직업들의 정보를 담고 있는 스트럭처
type role struct {
	// 각 직업들의 진영
	// 0: 무두장이, 1: 마을주민, 2: 늑대인간
	faction int

	// 각 직업별 행동 함수를 다르게 정의하기 위한 함수 선언
	action func(uid1, uid2 string, disRoleIdx int, player *user, g *game)
	String func() string
}

// 각 직업의 간단한 소개를 출력하는 함수.
func (r *role) getRoleInfo() string {
	return getRoleInfoFromDB(r.String())
}

// 유저의 직업을 반환
func (r *role) getRole(uid string, g *game) role {
	// TODO 내부 구현.
}

// 두 유저의 직업을 서로 교환
func (r *role) swapRole(uid1, uid2 string, g *game) {
	// TODO 내부 구현.
}

// 유저 직업과 버려진 직업을 교환.
func (r *role) switchRole(uid string, disRoleIdx int, g *game) {
	// TODO 내부 구현.
}

// 버려진 직업 중 하나 확인.
func (r *role) getDiscard(disRoleIdx int, g *game) role {
	// TODO 내부 구현.
}

// 특정 직업의 유저 목록 반환.
func (r *role) getRoleUser(find *role, g *game) (users []user) {
	// TODO 내부 구현.
}

// 모든 사람들의 직업을 입장순서별로 한칸 회전.
func (r *role) rotateRole(g *game) {
	// TODO 내부 구현.
}

// 유저에게 특수권한 부여
func (r *role) givePower(uid string, g *game) {
	// TODO 내부 구현.
}

// 특정 유저의 직업을 복사.
func (r *role) copyRole(destUID, srcUID string) {
	// TODO 내부 구현.
}
