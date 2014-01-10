package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"time"
)

var (
	CFG_PATH           = os.Getenv("HOME") + "/.file-server/config.json"
	BaseDir            string
	CacheExpireMinutes int64
	MaxProcs           int
	ServerAddr         string
	DebugAddr          string
	ReadTimeout        time.Duration
	WriteTimeout       time.Duration
	MainLogFile        string
	AccessLogFile      string
)

func init() {
	if data, err := ioutil.ReadFile(CFG_PATH); err != nil {
		log.Fatal(err)
	} else {
		var cfg map[string]interface{}
		if err = json.Unmarshal(data, &cfg); err != nil {
			log.Fatal(err)
		} else {
			BaseDir = cfg["base_dir"].(string)
			CacheExpireMinutes = int64(cfg["cache_expire_minutes"].(float64))
			MaxProcs = int(cfg["max_procs"].(float64))
			ServerAddr = cfg["server_addr"].(string)
			DebugAddr = cfg["debug_addr"].(string)
			ReadTimeout = time.Duration(cfg["read_timeout"].(float64)) * time.Second
			WriteTimeout = time.Duration(cfg["write_timeout"].(float64)) * time.Second
			MainLogFile = cfg["main_log"].(string)
			AccessLogFile = cfg["access_log"].(string)
		}
	}
}
