package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func mongoConn() (client *mongo.Client, ctx context.Context) {
	// timeout 기반의 Context 생성
	ctx, _ = context.WithTimeout(context.Background(), time.Second*4)

	// Authetication 을 위한 Client Option 구성
	clientOptions := options.Client().ApplyURI(
		env["dbURI"]).SetAuth(
		options.Credential{
			AuthSource: "",
			Username:   env["dbUserName"],
			Password:   env["dbPassword"],
		},
	)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		loggerError.Println("MongoDB connect Fail")
		loggerDebug.Println("MongoDB connect Fail")
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		loggerError.Println("MongoDB client ping Fail")
		loggerDebug.Println("MongoDB client ping Fail")
		log.Fatal(err)
	}
	loggerError.Println("MongoDB Connection Success")
	loggerDebug.Println("MongoDB Connection Success")
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

func roleGuideInsert(mongoDB *mongo.Database) {
	jsonFile, err := os.Open("Asset/role_guide.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	var byteValue []byte
	byteValue, err = ioutil.ReadAll(jsonFile)
	if err != nil {
		loggerError.Println("json file open error: ", err)
		log.Fatal(err)
	}
	var bDoc bson.D
	bson.Unmarshal([]byte(byteValue), &bDoc)
	col := mongoDB.Collection("Asset.role_guide")

	insRes, err := col.InsertMany(context.TODO(), byteValue)
	if err != nil {
		loggerError.Println("MongoDB Insert error: ", bDoc)
		log.Fatal(err)
	}
	loggerDebug.Printf("Insert role_guide SUCCESS: Inserted %v documents\n", insRes.InsertedCount)

	if err != nil {
		loggerError.Println("Collection load error: ", err)
		log.Fatal(err)
	}
	delRes, err := col.DeleteMany(context.TODO(), bson.D)
	if err != nil {
		loggerError.Println("MongoDB Delete error: ", err)
		log.Fatal(err)
	}
	loggerDebug.Printf("Delete role_guide SUCCESS: Deleted %v documents\n", delRes.DeletedCount)

	insRes, err = col.InsertMany(context.TODO(), byteValue)
	if err != nil {
		loggerError.Println("MongoDB Insert error: ", bDoc)
		log.Fatal(err)
	}
	loggerDebug.Printf("Insert role_guide SUCCESS: Inserted %v documents\n", insRes.InsertedCount)
}
