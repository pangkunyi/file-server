package logs

import (
	"config"
	"log"
	"os"
)

var (
	MLog *log.Logger
	ALog *log.Logger
)

func init() {
	mainLogFile, err := os.OpenFile(config.MainLogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	MLog = log.New(mainLogFile, "", log.LstdFlags)
	accessLogFile, err := os.OpenFile(config.AccessLogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	ALog = log.New(accessLogFile, "", log.LstdFlags)
}
