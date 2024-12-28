package config

import (
	"log"

	"gopkg.in/ini.v1"
)



type ConfigList struct {
	Port      string
	SQLDriver string
	DbName    string
	LogFile   string
}


var Config ConfigList

// main関数より先に読み込みたい用
func init(){
	LoadConfig()
}

func LoadConfig(){
	cfg,err := ini.Load("todo_app/config/config.ini")
	if err != nil {
		log.Fatalln(err)
	}

	Config = ConfigList{
		Port:      cfg.Section("web").Key("port").MustString("8080"),
		SQLDriver: cfg.Section("db").Key("driver").String(),  // DBセクションから読み込む
		DbName:    cfg.Section("db").Key("name").String(),    // DBセクションから読み込む
		LogFile:   cfg.Section("web").Key("logfile").String(),
	}
}
