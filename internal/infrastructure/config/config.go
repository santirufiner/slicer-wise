package config

import (
	"github.com/kelseyhightower/envconfig"
	"time"
)

type Config struct {
	LogLevel string        `required:"true" envconfig:"LOG_LEVEL"`
	Port     int           `required:"true" envconfig:"PORT"`
	Timeout  time.Duration `required:"false" envconfig:"BASE_HTTP_TIMEOUT" default:"40s"`
	Pg       Pg            `required:"true"`
}
type Pg struct {
	Url       string        `required:"true" envconfig:"PG_URL"`
	Timeout   time.Duration `required:"true" envconfig:"PG_TIMEOUT"`
	Heartbeat time.Duration `envconfig:"PG_HEARTBEAT" default:"10s"`

	RunMigration    bool   `envconfig:"RUN_MIGRATION" default:"false"`
	SourceMigration string `envconfig:"SOURCE_MIGRATION" default:"./scripts/migrations"`
}

func Load() Config {
	var config Config
	envconfig.MustProcess("", &config)
	return config
}
