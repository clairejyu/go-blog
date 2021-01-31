package db

import (
	"flag"
	"log"

	"github.com/go-ini/ini"
)

var cfg *ini.File

type DatabaseConfig struct {
	Type     string
	User     string
	Password string
	Host     string
	Port     string
	Name     string
}

var DBConfig = &DatabaseConfig{}

func Setup() {
	var err error

	if flag.Lookup("test.v") == nil {
		cfg, err = ini.Load("/Users/jyu/go/src/github.com/clairejyu/go-blog/config/config_dev.ini")
	} else {
		cfg, err = ini.Load("/Users/jyu/go/src/github.com/clairejyu/go-blog/config/config_test.ini")
	}

	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'config.ini': %v", err)
	}

	mapTo("database", DBConfig)
}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
