package config

import (
	"log"

	"github.com/go-ini/ini"
)

var (
	ServerConf = &ServerConfig{}
	MySqlConf  = &MySqlConfig{}
	WxConf     = &WxConf{}
)

type ServerConfig struct {
	Port     string `toml:"port"`
	RunMode  string `toml:"run-mode"`
	FilePath string `toml:"file-path"`
}

type MySqlConfig struct {
	Host     string `toml:"host"`
	Port     string `toml:"port"`
	DbName   string `toml:"dbname"`
	UserName string `toml:"username"`
	PassWd   string `toml:"passwd"`
}

type WxConfig struct {
	AppId     string `toml:"appId"`
	AppSecret string `toml:"appSecret"`
}

func init() {
	log.Println("loading app.conf")

	cfg, err := ini.Load("conf/app.conf")
	if err != nil {
		log.Fatal("load app conf err:%v", err)
	}
	err = cfg.Section("server").MapTo(ServerConf)
	if err != nil {
		log.Fatal("init server conf err:", err)
	}
	err = cfg.Section("mysql").MapTo(MySqlConf)
	if err != nil {
		log.Fatal("init mysql conf err:", err)
	}
	err = cfg.Section("wx").MapTo(WxConf)
	if err != nil {
		log.Fatal("init wx conf err:", err)
	}
}
