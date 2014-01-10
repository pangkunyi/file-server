package logs

import (
	"config"
	"log"
	"os"
)

var (
	ALog *log.Logger
)

func init() {
	mainLogFile, err := os.OpenFile(config.MainLogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(mainLogFile)
	accessLogFile, err := os.OpenFile(config.AccessLogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	ALog = log.New(accessLogFile, "", log.LstdFlags)
}
