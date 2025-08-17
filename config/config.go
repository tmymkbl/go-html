package config

import (
	"go-sample-todo/utils"
	"log"

	"gopkg.in/go-ini/ini.v1"
)

type ConfigList struct {
	Port      string
	SQLDriver string
	DbName    string
	DbUser    string
	DbPasswd  string
	DbHost    string
	DbPort    string
	Static    string
	Assets    string
	AdminURL  string
	LogOutput string
	LogFile   string
	LogLevel  string
}

var Config ConfigList

func init() {
	LoadConfig()
	utils.LoggingSettings(Config.LogFile)
}

func LoadConfig() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalln(err)
	}
	Config = ConfigList{
		Port:      cfg.Section("web").Key("port").MustString("8880"),
		SQLDriver: cfg.Section("db").Key("driver").String(),
		DbName:    cfg.Section("db").Key("name").String(),
		DbUser:    cfg.Section("db").Key("user").String(),
		DbPasswd:  cfg.Section("db").Key("passwd").String(),
		DbHost:    cfg.Section("db").Key("host").String(),
		DbPort:    cfg.Section("db").Key("port").String(),
		Static:    cfg.Section("web").Key("static").String(),
		Assets:    cfg.Section("web").Key("assets").String(),
		AdminURL:  cfg.Section("web").Key("admin_url").String(),
		LogFile:   cfg.Section("logging").Key("file").MustString(""),
		LogLevel:  cfg.Section("logging").Key("level").MustString("info"),
		LogOutput: cfg.Section("logging").Key("output").MustString("stdout"),
	}
}
