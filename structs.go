package main

import (
	"math/rand"
	"time"

	"github.com/bwmarrin/discordgo"
	embed "github.com/clinet/discordgo-embed"
)

type Guild_Info struct {
	GID          string
	CID          string
	IngameUserID []string
	NumOfPlayer  int
	CardDeck     []string
	RoleUID      map[string]string
	GameLog      []string
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
