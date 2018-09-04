package config

import (
	"github.com/7phs/coding-challenge-search/helper"
	log "github.com/sirupsen/logrus"
	"gopkg.in/go-playground/validator.v9"
)

const (
	EnvConfigAddr        = "ADDR"
	EnvConfigCors        = "CORS"
	EnvConfigDatabaseUrl = "DB_URL"

	DefaultAddr = ":8080"
	DefaultCors = false
)

var (
	Conf *Config
)

type Config struct {
	Addr        string `validate:"required, tcp_addr"`
	Cors        bool
	DatabaseUrl string `validate:"required"`
}

func NewConfig() *Config {
	return &Config{
		Addr:        helper.GetEnvStr(EnvConfigAddr, DefaultAddr),
		Cors:        helper.GetEnvBool(EnvConfigCors, DefaultCors),
		DatabaseUrl: helper.GetEnvStr(EnvConfigDatabaseUrl, ""),
	}
}

func (o *Config) Validate() {
	validate := validator.New()

	err := validate.Struct(o)
	if err != nil {
		log.Errorf("config: %+v", err)
	}
}

func Init() {
	Conf = NewConfig()

	Conf.Validate()
}
