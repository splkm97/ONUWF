package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func loadFile(path string) string {
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(dat)
}

func main() {
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
		panic(err)
	}
	fmt.Println("Connect Success")
	err = client.Ping(context.Background(), nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Ping Success")
}
