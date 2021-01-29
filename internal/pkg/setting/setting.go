package setting

import (
	"log"

	"github.com/go-ini/ini"
)

var cfg *ini.File

type Database struct {
	Type     string
	User     string
	Password string
	Host     string
	Port     string
	Name     string
}

var DatabaseSetting = &Database{}

func Setup() {
	var err error
	cfg, err = ini.Load("/Users/jyu/go/src/github.com/clairejyu/go-blog/config/config_dev.ini")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'config/config_dev.ini': %v", err)
	}

	mapTo("database", DatabaseSetting)
}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
