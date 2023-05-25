package options

import (
	"encoding/json"
	pkg_options "gin_test/internal/pkg/options"
	"gin_test/pkg/log"
)

type Options struct {
	PostgresConfig *pkg_options.PostgresOptions `json:"postgres"`
	RedisConfig    *pkg_options.RedisOptions    `json:"redis"`
	Log            *log.Options                 `json:"log"`
}

// NewOptions create a new Options obj with default params
func NewOptions() *Options {
	o := Options{
		PostgresConfig: pkg_options.NewPostgresOptions(),
		RedisConfig:    pkg_options.NewRedisOptions(),
		Log:            log.NewOptions(),
	}

	return &o
}

func (o *Options) String() string {
	data, _ := json.Marshal(o)
	return string(data)
}
