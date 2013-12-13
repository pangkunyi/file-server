package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

var (
	CFG_PATH           = os.Getenv("HOME") + "/.file-server/config.json"
	BaseDir            string
	CacheExpireMinutes int64
	MaxProcs           int
	ServerAddr         string
)

func init() {
	if data, err := ioutil.ReadFile(CFG_PATH); err != nil {
		panic(err)
	} else {
		var cfg map[string]interface{}
		if err = json.Unmarshal(data, &cfg); err != nil {
			panic(err)
		} else {
			BaseDir = cfg["base_dir"].(string)
			CacheExpireMinutes = int64(cfg["cache_expire_minutes"].(float64))
			MaxProcs = int(cfg["max_procs"].(float64))
			ServerAddr = cfg["server_addr"].(string)
		}
	}
}
