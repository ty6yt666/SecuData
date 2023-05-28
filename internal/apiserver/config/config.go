package config

import "gin_test/internal/apiserver/options"

type Config struct {
	*options.Options
}

type Res struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}
