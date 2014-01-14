package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"time"
)

var (
	CFG_PATH = os.Getenv("HOME") + "/.file-server/config.json"
	C        Config
)

type Config struct {
	CacheExpireMinutes int64         `json:"cache_expire_minutes"`
	MaxProcs           int           `json:"max_procs"`
	ServerAddr         string        `json:"server_addr"`
	DebugAddr          string        `json:"debug_addr"`
	ReadTimeout        time.Duration `json:"read_timeout"`
	WriteTimeout       time.Duration `json:"write_timeout"`
	MainLogFile        string        `json:"main_log"`
	AccessLogFile      string        `json:"access_log"`
	Rules              []Rule        `json:"rules"`
}

type Rule struct {
	Cached  bool   `json:"cached"`
	Pattern string `json:"pattern"`
	Strip   string `json:"strip"`
	Dir     string `json:"dir"`
}

func init() {
	if data, err := ioutil.ReadFile(CFG_PATH); err != nil {
		log.Fatal(err)
	} else {
		if err = json.Unmarshal(data, &C); err != nil {
			log.Fatal(err)
		}
		C.ReadTimeout = C.ReadTimeout * time.Second
		C.WriteTimeout = C.WriteTimeout * time.Second
	}
}
