package config

import (
	"github.com/7phs/coding-challenge-search/helper"
	"github.com/c2h5oh/datasize"
	log "github.com/sirupsen/logrus"
	"gopkg.in/go-playground/validator.v9"
)

const (
	EnvConfigAddr          = "ADDR"
	EnvConfigCors          = "CORS"
	EnvConfigStage         = "STAGE"
	EnvConfigDatabaseUrl   = "DB_URL"
	EnvConfigKeywordsLimit = "KEYWORDS_LIMIT"

	DefaultAddr          = ":8080"
	DefaultCors          = false
	DefaultStage         = "development"
	DefaultKeywordsLimit = int64(4 * datasize.KB)
)

var (
	Conf *Config
)

func Init() {
	Conf = NewConfig()

	Conf.Validate()

	Conf.LogInfo()
}

type Config struct {
	Addr          string `validate:"required"`
	Cors          bool
	Stage         string
	DatabaseUrl   string `validate:"required"`
	KeywordsLimit datasize.ByteSize
}

func NewConfig() *Config {
	return &Config{
		Addr:          helper.GetEnvStr(EnvConfigAddr, DefaultAddr),
		Cors:          helper.GetEnvBool(EnvConfigCors, DefaultCors),
		Stage:         helper.GetEnvStr(EnvConfigStage, DefaultStage),
		DatabaseUrl:   helper.GetEnvStr(EnvConfigDatabaseUrl, "./fatlama.sqlite3"),
		KeywordsLimit: datasize.ByteSize(helper.GetEnvInt64(EnvConfigKeywordsLimit, DefaultKeywordsLimit)),
	}
}

func (o *Config) Validate() {
	validate := validator.New()

	err := validate.Struct(o)
	if err != nil {
		log.Errorf("config: %+v", err)
	}
}

func (o *Config) LogInfo() {
	for _, info := range []*struct {
		name  string
		value interface{}
	}{
		{name: "addr", value: o.Addr},
		{name: "cors", value: o.Cors},
		{name: "stage", value: o.Stage},
		{name: "db_url", value: o.DatabaseUrl},
		{name: "keywords_limit", value: o.KeywordsLimit},
	} {
		log.Info("config: "+info.name+" - '", info.value, "'")
	}
}
