package config

import (
	"api/utils"
)

type HttpConfigInterface struct {
	Domain             string   `toml:"domain"`
	Port               int      `toml:"port"`
	Version            string   `toml:"version"`
	SessionLiveSecond  int      `toml:"session-live-second"`
	SessionLiveMaxHour int      `toml:"session-live-max-hour"`
	TrustProxyIpCidr   []string `toml:"trust-proxy-ip-cidr"`
}

type DBConfigInterface struct {
	DataSource string `toml:"dataSource"`
	MaxIdle    int    `toml:"max-idle-connections"`
	MaxOpen    int    `toml:"max-open-connections"`
}

type RedisConfigInterface struct {
	Addr     string `toml:"addr"`
	Password string `toml:"password"`
}

type ConfInterface struct {
	Salt string `toml:"salt"`
}

type CombinationConfig struct {
	Http  HttpConfigInterface  `toml:"http"`
	DB    DBConfigInterface    `toml:"mysql"`
	Redis RedisConfigInterface `toml:"redis"`
	Conf  ConfInterface        `toml:"conf"`
}

var DefaultConfig = CombinationConfig{
	Http: HttpConfigInterface{
		Domain:             "http://localhost/",
		Port:               8090,
		Version:            "1.0.0",
		SessionLiveSecond:  14400,
		SessionLiveMaxHour: 24,
	},
	DB: DBConfigInterface{
		DataSource: "user:password@tcp(127.0.0.1:3306)/dbname?charset=utf8",
		MaxIdle:    20,
		MaxOpen:    500,
	},
	Redis: RedisConfigInterface{
		Addr:     "localhost:6379",
		Password: "",
	},
	Conf: ConfInterface{Salt: utils.RandomStr(24)},
}
