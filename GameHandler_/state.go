package main

import (
	"github.com/bwmarrin/discordgo"
)

type state interface {
	// 사용자 인원수 3 ~ 26
	// num: 0 ~ 23
	// pressNumBtn 사용자가 숫자 이모티콘을 눌렀을 때 state에서 하는 동작
	pressNumBtn(s *discordgo.Session, r *discordgo.MessageReactionAdd, num int)

	// pressDisBtn 사용자가 버려진 카드 이모티콘을 눌렀을 때 state에서 하는 동작
	pressDisBtn(s *discordgo.Session, r *discordgo.MessageReactionAdd)

	// pressYesBtn 사용자가 yes 이모티콘을 눌렀을 때 state에서 하는 동작
	pressYesBtn(s *discordgo.Session, r *discordgo.MessageReactionAdd)

	// pressNoBtn 사용자가 No 이모티콘을 눌렀을 때 state에서 하는 동작
	pressNoBtn(s *discordgo.Session, r *discordgo.MessageReactionAdd)

	// pressDirBtn 좌 -1, 우 1 사용자가 좌우 방향 이모티콘을 눌렀을 때 state에서 하는 동작
	pressDirBtn(s *discordgo.Session, r *discordgo.MessageReactionAdd, dir int)
}
