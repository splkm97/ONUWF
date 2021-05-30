package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

var (
	env map[string]string
	rg  []roleGuide
	emj map[string]string

	loggerLog   *log.Logger
	loggerError *log.Logger
	loggerUser  *log.Logger
	loggerDebug *log.Logger
)

type roleGuide struct {
	roleName  string `json:"roleName"`
	roleGuide string `json:"roleGuide"`
	max       int    `json:"max"`
	faction   string `json:"faction"`
}

// 설치 환경 불러오기.
func envInit() {
	envFile, err := os.Open("env.json")
	defer envFile.Close()
	if err != nil {
		log.Fatal(err)
	}
	defer envFile.Close()

	var byteValue []byte
	byteValue, err = ioutil.ReadAll(envFile)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal([]byte(byteValue), &env)
}

// 직업 가이드 에셋 불러오기.
func roleGuideInit() {
	rgFile, err := os.Open("Asset/role_guide.json")
	defer rgFile.Close()
	if err != nil {
		log.Fatal(err)
	}
	defer rgFile.Close()

	var byteValue []byte
	byteValue, err = ioutil.ReadAll(rgFile)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal([]byte(byteValue), &rg)
}

// 이모지 맵에 불러오기.
func emojiInit() {
	emjFile, err := os.Open("Asset/emoji.json")
	if err != nil {
		log.Fatal(err)
	}
	defer emjFile.Close()

	var byteValue []byte
	byteValue, err = ioutil.ReadAll(emjFile)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal([]byte(byteValue), &emj)
}

// 로거 변수 초기화.
func loggerInit() {
	logErrorFile, err := os.OpenFile(env["logErrorPath"], os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	loggerLog = log.New(logErrorFile, "LOG: ", log.Ldate|log.Ltime|log.Lshortfile)
	loggerError = log.New(logErrorFile, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	var logUserFile *os.File
	logUserFile, err = os.OpenFile(env["logUserPath"], os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		loggerError.Println("Can not open env['logUserPath']:", env["logUserPath"])
		log.Fatal(err)
	}
	loggerUser = log.New(logUserFile, "USER: ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)

	var logDebugFile *os.File
	logDebugFile, err = os.OpenFile(env["logDebugPath"], os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		loggerError.Println("Can not open env['logUserPath']:", env["logUserPath"])
		log.Fatal(err)
	}
	loggerDebug = log.New(logDebugFile, "DEBUG: ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
}
