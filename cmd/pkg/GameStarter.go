// +build linux,amd64,go1.15,!cgo

package pkg

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"

	gh "../internal/gamehandler"
	"github.com/bwmarrin/discordgo"
)

var (
	isInGame map[string]bool
)

func init() {
	gh.EnvInit()
	gh.RoleGuideInit()
	gh.LoggerInit()

	isInGame = make(map[string]bool)
	/*
		conn, ctx := mongoConn()
		mongoDB := conn.Database("WF_Data")

		data := allData("people", mongoDB, ctx)
		fmt.Println(string(data))
	*/
}

func main() {
	dg, err := discordgo.New("Bot " + env["dgToken"])
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}
	dg.AddHandler(messageCreate)
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	dg.Close()
}

func startGame(s *discordgo.Session, m *discordgo.MessageCreate) {
	path, _ := exec.LookPath("./GameHandler_/GameHandler")
	args := make([]string, 3)
	args[0] = "-gid=" + m.GuildID
	args[1] = "-cid=" + m.ChannelID
	args[2] = "-mid=" + m.Author.ID
	gameHandlerCmd := exec.Command(path, args...)
	err := gameHandlerCmd.Run()

	if err != nil {
		msg := "게임 핸들러를 실행하지 못하였습니다. : "
		guild, _ := s.Guild(m.GuildID)
		channel, _ := s.Channel(m.ChannelID)
		msg += "Guild: " + guild.Name
		msg += ", Channel: " + channel.Name
		msg += ", Master: " + m.Author.Username
		msg += ", error: "
		loggerError.Println(msg, err)
		return
	}
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "ㅁ시작" {
		if isInGame[m.GuildID+m.ChannelID] {
			s.ChannelMessageSend(m.ChannelID, "현재 게임이 진행중인 채널입니다.")
			return
		}
		isInGame[m.GuildID+m.ChannelID] = true
		go startGame(s, m)
	}
	if m.Content == "ㅁ강제종료" {
		if isInGame[m.GuildID+m.ChannelID] {
			isInGame[m.GuildID+m.ChannelID] = false
		}
	}
}
