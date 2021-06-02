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
)

type roleGuide struct {
	RoleName  string   `json:"roleName"`
	RoleGuide []string `json:"roleGuide"`
	Max       int      `json:"max"`
	Faction   string   `json:"faction"`
}

// 설치 환경 불러오기.
func envInit() {
	envFile, err := os.Open("env.json")
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
