package main

// role werewolf struct
type struct werewolf {
	// role을 상속받는 구조체.
	role
}

// 늑대인간을 반환하는 함수.
func (wf *werewolf) String() string {
	return "늑대인간"
}

// 늑대인간의 능력을 사용하는 함수.
func (wf *werewolf) Action(uid1, uid2 string, disRole int, g *game) {
	
}
