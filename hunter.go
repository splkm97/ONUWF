package main

// role hunter struct
type hunter struct {
	role
}

func (ht *hunter) String() string {
	return "사냥꾼"
}

func (ht *hunter) Action(...) {
	ht.givePower(HUNTER_POW, g)
}
