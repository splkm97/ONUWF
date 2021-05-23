package main

/*
var testGuild Guild_Info

func testinit() {
	testGuild = Guild_Info{
		GID: "1234",
		CID: "1234",
		IngameUserID: []string{
			"1234",
			"1235",
		},
		NumOfPlayer: 2,
		CardDeck:    nil,
		RoleUID:     nil,
		GameLog:     nil,
	}
}

func TestGenLogMsg(t *testing.T) {
	testinit()
	testGuild.AddGameLog("this")
	testGuild.AddGameLog("is")
	testGuild.AddGameLog("game")
	testGuild.AddGameLog("log")
	got := genLogMsg(&testGuild)
	want := "\nthis\nis\ngame\nlog"

	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestCardSeqGen(t *testing.T) {
	t.Run("6 \"test\" string input", func(t *testing.T) {
		testinit()
		for i := 0; i < 6; i++ {
			testGuild.AddCardDeck("test")
		}
		nextCard := testGuild.CardSeqGen()
		for i := 0; i < 6; i++ {
			got := nextCard()
			want := "test"
			if got != want {
				t.Errorf("i: %v, got: %v, want: %v", i, got, want)
			}
		}
	})

	t.Run("wolf 2, vil 2, rob 1", func(t *testing.T) {
		testinit()
		testGuild.AddCardDeck("wolf")
		testGuild.AddCardDeck("wolf")
		testGuild.AddCardDeck("vil")
		testGuild.AddCardDeck("vil")
		testGuild.AddCardDeck("rob")
		nextCard := testGuild.CardSeqGen()
		got := make([]string, 0)
		for i := 0; i < 5; i++ {
			got = append(got, nextCard())
		}
		want := make([]string, len(testGuild.CardDeck))
		copy(want, testGuild.CardDeck)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got:\n%v,\nwant:\n%v", got, want)
		}
	})

	t.Run("shuffle and print", func(t *testing.T) {
		testinit()
		testGuild.AddCardDeck("wolf")
		testGuild.AddCardDeck("wolf")
		testGuild.AddCardDeck("vil")
		testGuild.AddCardDeck("rob")
		testGuild.AddCardDeck("ora")
		testGuild.AddCardDeck("dru")
		testGuild.AddCardDeck("tan")
		testGuild.AddCardDeck("min")
		nextOriCard := testGuild.CardSeqGen()
		testGuild.ShuffleDeck()
		nextCard := testGuild.CardSeqGen()

		var cnt int
		for i := 0; i < 8; i++ {
			nxt := nextCard()
			if fmt.Println(nxt); nextOriCard() == nxt {
				cnt++
			}
		}
		if cnt == 6 {
			t.Errorf("all card seq is same")
		}
	})
}

*/
