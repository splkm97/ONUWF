package main

import "github.com/bwmarrin/discordgo"

// role hunter struct
type hunter struct {
	role
}

func (ht *hunter) String() string {
	return "사냥꾼"
}

func (ht *hunter) Action(uid1, uid2 string, disRole int, player *user, g *game, s *discordgo.Session) {
	ht.givePower(HUNTER_POW, g)
}
