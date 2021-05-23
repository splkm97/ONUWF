package main

import (
	"os"
	"log"
	"io/ioutil"
)

var (
	env			map[string] string

	loggerLog	*log.Logger
	loggerError	*log.Logger
	loggerUser	*log.Logger
	loggerDebug	*log.Logger
)

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

func loggerInit() {
	logErrorFile, err := os.OpenFile(env["logErrorPath"], os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer logErrorFile.Close()
	loggerLog = log.New(logErrorFile, "LOG: ", log.Ldate|log.Ltime|log.Lshortfile)
	loggerError = log.New(logErrorFile, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	var	logUserFile *File
	logUserFile, err = os.OpenFile(env["logUserPath"], os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		LoggerError.Println("Can not open env['logUserPath']:", env["logUserPath"])
		log.Fatal(err)
	}
	defer logUserFile.Close()
	loggerUser = log.New(logUserFile, "USER: ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)

	var	logDebugFile *File
	logDebugFile, err = os.OpenFile(env["logDebugPath"], os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		LoggerError.Println("Can not open env['logUserPath']:", env["logUserPath"])
		log.Fatal(err)
	}
	defer logDebugFile.Close()
	loggerDebug = log.New(logDebugFile, "DEBUG: ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
}
