package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Token string
	mongo *mongo.Database
)

func loadFile(path string) string {
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return string(dat)
}

func init() {
	myStrToken := loadFile("./Auth/Token")
	conn := mongoConn()
	mongo := conn.Database("WF_Data")
	flag.StringVar(&Token, "t", myStrToken, "Bot Token")
	flag.Parse()
}

func mongoConn() (client *mongo.Client) {
	dbPath := "./Auth/Passwd"
	dbPwd := loadFile(dbPath)

	// timeout 기반의 Context 생성
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(10000000000))
	// Authetication 을 위한 Client Option 구성
	clientOptions := options.Client().ApplyURI(
		"mongodb://localhost:49154").SetAuth(
		options.Credential{
			AuthSource: "",
			Username:   "kalee",
			Password:   dbPwd,
		},
	)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB Connection Success")
	return client
}

func main() {
	dg, err := discordgo.New("Bot " + Token)
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

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == m.State.User.ID {
		return
	}
	if m.Content == "ping" {
		m.ChannelMessageSend(m.ChannelID, "Pong!")
	}
	if m.Content == "pong" {
		m.ChannelMessageSend(m.ChannelID, "Ping!")
	}
}
