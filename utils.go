package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func mongoConn() (client *mongo.Client, ctx context.Context) {
	dbPath := "./Auth/Passwd"
	dbPwd := loadFile(dbPath)

	// timeout 기반의 Context 생성
	ctx, _ = context.WithTimeout(context.Background(), time.Duration(500000000))
	// Authetication 을 위한 Client Option 구성
	clientOptions := options.Client().ApplyURI(
		"mongodb://localhost:49153").SetAuth(
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
	return client, ctx
}

func GetCollection(colName string, mongoDB *mongo.Database) *mongo.Collection {
	return mongoDB.Collection(colName)
}

func AllData(collection string, mongoDB *mongo.Database, ctx context.Context) string {
	// 데이터를 담을 변수 선언
	var datas []bson.M

	// 데이터 읽기
	res, err := GetCollection(collection, mongoDB).Find(ctx, bson.M{})

	// 결과를 변수에 담기
	if err = res.All(ctx, &datas); err != nil {
		fmt.Println(err)
	}

	// []byte를 String타입으로 변환
	data, _ := json.MarshalIndent(&datas, "", "	")

	return string(data)
}

func loadFile(path string) string {
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return string(dat)
}
