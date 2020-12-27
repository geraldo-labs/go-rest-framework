package settings

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)

const (
	DatabaseTypeSqlite   = "sqlite"
	DatabaseTypePostgres = "postgres"
	DatabaseTypeMysql    = "mysql"
)

type Settings struct {
	DatabaseType string `envconfig:"DATABASE_TYPE" default:"sqlite"`
	DatabaseDSN  string `envconfig:"DATABASE_DSN" default:"sqlite.db"`
}

func (s *Settings) Validate() error {
	return nil
}

func NewSettings(prefix string) *Settings {
	conf := &Settings{}
	if err := envconfig.Process(prefix, conf); err != nil {
		log.Fatal(err)
	}
	return conf
}
