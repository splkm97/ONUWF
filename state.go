package main

import (
	"github.com/bwmarrin/discordgo"
)

type state interface
{
	//사용자 인원수 3 ~ 26
	//num: 0 ~ 23
	//사용자가 숫자 이모티콘을 눌렀을 때 state에서 하는 동작
	pressNumBtn func(s *discordgo.Session, r *discordgo.MessageReactionAdd, num int)

	//사용자가 버려진 카드 이모티콘을 눌렀을 때 state에서 하는 동작
	pressDisBtn func(s *discordgo.Session, r *discordgo.MessageReactionAdd)

	//사용자가 yes 이모티콘을 눌렀을 때 state에서 하는 동작
	pressYesBtn func(s *discordgo.Session, r *discordgo.MessageReactionAdd)

	//사용자가 No 이모티콘을 눌렀을 때 state에서 하는 동작
	pressNoBtn func(s *discordgo.Session, r *discordgo.MessageReactionAdd)

	//좌 -1, 우 1
	//사용자가 좌우 방향 이모티콘을 눌렀을 때 state에서 하는 동작
	pressDirBtn func(s *discordgo.Session, r *discordgo.MessageReactionAdd, dir int)

	//사용자가 종료 메세지를 보냈을 때 state에서 하는 동작
	sendFinish func(s *discordgo.Session, m *discordgo.MessageCreate)
}

type struct StatePrepare
{
	//각 상태에서 저장 중인 게임
	g	*game
}

//사용자가 숫자 이모티콘을 눌렀을 때 StatePrepare에서 하는 동작
func (sPrepare *StatePrepare) pressNumBtn (s *discordgo.Session, r *discordgo.MessageReactionAdd, num int)
{
	sPrepare.game
}

//사용자가 버려진 카드 이모티콘을 눌렀을 때 StatePrepare에서 하는 동작
func (sPrepare *StatePrepare) pressDisBtn (s *discordgo.Session, r *discordgo.MessageReactionAdd)
{

}

//사용자가 yes 이모티콘을 눌렀을 때 StatePrepare에서 하는 동작
func (sPrepare *StatePrepare) pressDisBtn (s *discordgo.Session, r *discordgo.MessageReactionAdd)
{

}

//사용자가 No 이모티콘을 눌렀을 때 StatePrepare에서 하는 동작
func (sPrepare *StatePrepare) pressDisBtn (s *discordgo.Session, r *discordgo.MessageReactionAdd)
{

}

//좌 -1, 우 1
//사용자가 좌우 방향 이모티콘을 눌렀을 때 StatePrepare에서 하는 동작
func (sPrepare *StatePrepare) pressDisBtn (s *discordgo.Session, r *discordgo.MessageReactionAdd, dir int)
{

}

//사용자가 종료 메세지를 보냈을 때 StatePrepare에서 하는 동작
func (sPrepare *StatePrepare) sendFinish (s *discordgo.Session, m *discordgo.MessageCreate)
{

}