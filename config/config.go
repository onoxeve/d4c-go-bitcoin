package config

import (
	"github.com/labstack/gommon/log"
	"gopkg.in/ini.v1"
	"os"
)

type ConfigList struct {
	ApiKey    string
	ApiSecret string
	LogFile   string
}

var Config ConfigList

func init() {
	cfg, err := ini.Load("config/config.ini")
	if err != nil {
		log.Printf("Failed to read file: %v", err)
		os.Exit(1)
	}

	Config = ConfigList{
		ApiKey:    cfg.Section("bitflyer").Key("api_key").String(),
		ApiSecret: cfg.Section("bitflyer").Key("api_secret").String(),
		LogFile:   cfg.Section("app").Key("log_file").String(),
	}
}
