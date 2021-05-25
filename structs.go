package main

// 게임 진행을 위한 정보를 담고 있는 스트럭처
type GuildInfo struct {
	// 현재 게임이 진행중인 서버의 GID
	GID string

	// 현재 게임에서 설정된 직업들의 RID 목록
	RIDlist []int

	// 현재 게임의 참가자 UID
	UIDlist []string

	// 현재 게임의 진행시점
	CurStage string

	// Role을 User별로 매핑시킨 테이블
	// -1: 원래 직업, 0: 관계 없음, 1: 현재 직업
	RUidxTable [][]int

	// 게임 진행 상황을 기록하는 로그 메시지 목록
	LogMsg []string
}

// 각 직업들의 정보를 담고 있는 스트럭처
type Role struct {
	// 각 직업들의 고유한 RID
	RID int

	// 각 직업들의 이름
	Name string

	// 각 직업들의 능력순서
	Priority int

	// 각 직업들의 진영
	// 0: 무두장이, 1: 마을주민, 2: 늑대인간
	Faction int
}

// 각 사용자들의 정보를 담고 있는 스트럭처
type User struct {
	// 각 유저의 UID
	UID string

	// 각 유저가 속한 GID
	GID string

	// 각 유저의 RID
	RID int

	// 각 유저의 DM 채널 ID
	DMCID string
}

/*
type Game struct {
	ChannelID   string
	NumOfPlayer int
	UserIDRole  map[string]string
	Log         []string
}

func (g *Guild_Info) AddCardDeck(role string) {
	if g.CardDeck == nil {
		g.CardDeck = make([]string, 0)
	}
	g.CardDeck = append(g.CardDeck, role)
}

func (g *Guild_Info) ShuffleDeck() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(g.CardDeck), func(i, j int) {
		g.CardDeck[i], g.CardDeck[j] = g.CardDeck[j], g.CardDeck[i]
	})
}

func (g *Guild_Info) CardSeqGen() func() string {
	i := -1
	cardDeck := make([]string, len(g.CardDeck))
	copy(cardDeck, g.CardDeck)
	return func() string {
		i++
		if i == len(g.CardDeck) {
			return "ERROR_DECK_END"
		}
		return cardDeck[i]
	}
}

func (g *Guild_Info) AddGameLog(cont string) {
	if g.GameLog == nil {
		g.GameLog = make([]string, 0)
	}
	g.GameLog = append(g.GameLog, cont)
}

func genLogMsg(g *Guild_Info) string {
	var result string
	for _, item := range g.GameLog {
		result += "\n" + item
	}
	return result
}

func (g *Guild_Info) PrintGameLog(s *discordgo.Session) {
	logMsg := genLogMsg(g)
	s.ChannelMessageSendEmbed(
		g.CID,
		embed.NewGenericEmbed("게임 로그", logMsg),
	)
}
*/
