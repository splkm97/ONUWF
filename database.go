package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func mongoConn() (client *mongo.Client, ctx context.Context) {
	// timeout 기반의 Context 생성
	ctx, err := context.WithTimeout(context.Background(), time.Second*5)
	if err != nil {
		LoggerError.Println("MongoDB make context Fail")
		LoggerDebug.Println("MongoDB make context Fail")
		log.Fatal(err)
	}
	// Authetication 을 위한 Client Option 구성
	clientOptions := mongo.options.Client().ApplyURI(
		env["dbURI"]).SetAuth(
		mongo.options.Credential{
			AuthSource:	"",
			Username:	env["dbUserName"],
			Password:	env["dbPassword"],
		},
	)

	var client *Client
	client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		LoggerError.Println("MongoDB connect Fail")
		LoggerDebug.Println("MongoDB connect Fail")
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		LoggerError.Println("MongoDB client ping Fail")
		LoggerDebug.Println("MongoDB client ping Fail")
		log.Fatal(err)
	}
	LoggerError.Println("MongoDB Connection Success")
	LoggerDebug.Println("MongoDB Connection Success")
	log.Println("MongoDB Connection Success")
	return client, ctx
}

func allData(collection string, mongoDB *mongo.Database, ctx context.Context) string {
	var datas []bson.M
	res, err := mongoDB.Collection(collection).Find(ctx, bson.M{})
	if err = res.All(ctx, &datas); err != nil {
		fmt.Println(err)
	}

	data, _ := json.MarshalIndent(&datas, "", "	")

	return string(data)
}