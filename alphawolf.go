package main

import "github.com/bwmarrin/discordgo"

// role alphawolf struct
type alphawolf struct {
	// role 을 상속받는 구조체.
	role
}

// 직업명 '대장 늑대' 를 반환하는 함수.
func (aw *alphawolf) String() string {
	return "대장 늑대"
}

// 대장 늑대의 능력을 사용하는 함수.
func (aw *alphawolf) Action(uid1, uid2 string, disRole int, player *user, g *game, s *discordgo.Session) {
	// 늑대인간 구조체의 메소드 재활용.
	wfIns := werewolf{}
	wfIns.Action(uid1, uid2, disRole, player, g, s)
	if disRole == 4 {
		target := getRole(uid1, g)
		switchRole(uid1, disRole, g)
		msg := "대장 늑대 `" + player.nick + "` 은(는) "
		msg += "`" + target.String() + "` 였던 `" + g.findUserByUID(uid1).nick + "` 을(를)\n"
		msg += "늑대인간으로 변신시켰습니다."
		g.appendLog(msg)
	} else {
		g.appendLog("대장 늑대 `" + player.nick + "` 은(는) 능력을 사용하지 않았습니다.")
	}
}
