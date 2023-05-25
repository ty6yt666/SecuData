package options

import "encoding/json"

// RedisOptions defines options for redis cluster
type RedisOptions struct {
	Host                  string   `json:"host"`
	Port                  int      `json:"port"`
	Addrs                 []string `json:"addrs"`
	Username              string   `json:"username"`
	Password              string   `json:"password"`
	Database              int      `json:"database"`
	MasterName            string   `json:"master_name"`
	MaxIdle               int      `json:"optimisation_max_idle"`
	MaxActive             int      `json:"optimisation_max_active"`
	Timeout               int      `json:"timeout"`
	EnableCluster         bool     `json:"enable_cluster"`
	UseSSL                bool     `json:"use_ssl"`
	SSLInsecureSkipVerify bool     `json:"ssl_insecure_skip_verify"`
}

// NewRedisOptions create a default value instance
func NewRedisOptions() *RedisOptions {
	return &RedisOptions{
		Host:                  "127.0.0.1",
		Port:                  6379,
		Addrs:                 []string{},
		Username:              "",
		Password:              "",
		Database:              0,
		MasterName:            "",
		MaxIdle:               2000,
		MaxActive:             4000,
		Timeout:               0,
		EnableCluster:         false,
		UseSSL:                false,
		SSLInsecureSkipVerify: false,
	}
}

func (o *RedisOptions) String() string {
	data, _ := json.Marshal(o)
	return string(data)
}
