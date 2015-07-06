package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path"
	"time"
)

var (
	CFG_PATH = "/etc/" + path.Base(os.Args[0]) + "/config.json"
	C        Config
)

type Config struct {
	CacheExpireTime time.Duration `json:"cache_expire_time"`
	MaxProcs        int           `json:"max_procs"`
	ServerAddrs     []string      `json:"server_addrs"`
	DebugAddr       string        `json:"debug_addr"`
	ReadTimeout     time.Duration `json:"read_timeout"`
	WriteTimeout    time.Duration `json:"write_timeout"`
	MainLogFile     string        `json:"main_log"`
	AccessLogFile   string        `json:"access_log"`
	Rules           []Rule        `json:"rules"`
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
		C.CacheExpireTime = C.CacheExpireTime * time.Minute
		C.ReadTimeout = C.ReadTimeout * time.Second
		C.WriteTimeout = C.WriteTimeout * time.Second
	}
}
