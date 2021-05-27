package main

import "testing"

var (
	g1 game
)

func initGame() {
	g1.userlist = make([]user, 0)
	g1.userlist = append(g1.userlist, user{"001", "kalee", "chan1", "dmChan1"}, user{"002", "min-jo", "chan2", "dmChan2"}, user{"003", "juhur", "chan3", "dmChan3"})
	g1.rolelist = make([]role, 0)
	g1.rolelist = append(g1.rolelist, werewolf{}, werewolf{}, hunter{}, alphawolf{}, misticwolf{}, werewolf{})
	g1.roleIdxTable
}

func TestGetRole(t *testing.T) {
	initGame()
	rr := g1.getRole("001")
	got := rr.String()
	want := "늑대인간"
	if got != want {
		t.Errorf("got: %v, want: %v\n", got, want)
	}

}
