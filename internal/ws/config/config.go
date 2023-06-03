package config

import (
	"fmt"
	"gin_test/internal/apiserver/options"
	"github.com/go-ini/ini"
)

type Config struct {
	*options.Options
}

type Res struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}

var conf = new(Config)

func Init(file string) {
	cfg, err := ini.Load(file)
	if err != nil {
		recover()
	}
	cfg.MapTo(&conf)
	fmt.Println(cfg)
}
