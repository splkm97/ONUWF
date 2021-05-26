package main

import "github.com/bwmarrin/discordgo"

// role misticwolf struct
type misticwolf struct {
	role
}

// 직업명 '신비한 늑대' 를 반환하는 함수.
func (mw *misticwolf) String() string {
	return "신비한 늑대"
}

// 신비한 늑대 능력 사용 함수.
func (mw *misticwolf) Action(uid1, uid2 string, disRole int, player *user, g *game, s *discordgo.Session) {
	// 늑대인간과 예언자의 메소드 활용.
	wfIns := werewolf{}
	seerIns := seer{}
	wfIns.Action(uid1, uid2, disRole, player, g, s)
	seerIns.Action(uid, uid2, disRole, player, g, s)
}
