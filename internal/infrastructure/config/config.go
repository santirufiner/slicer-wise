package config

import (
	"github.com/kelseyhightower/envconfig"
	"time"
)

type Config struct {
	LogLevel string        `split_words:"true" required:"true"`
	Port     int           `required:"true"`
	Pg       Pg            `required:"true"`
	Timeout  time.Duration `default:"40s" envconfig:"BASE_HTTP_TIMEOUT"`
}
type Pg struct {
	Url       string        `required:"true" split_words:"true"`
	Timeout   time.Duration `required:"true"`
	Heartbeat time.Duration `default:"10s"`

	RunMigration    bool   `default:"false" split_words:"true"`
	SourceMigration string `default:"./scripts/migrations" split_words:"true"`
}

func Load() Config {
	var config Config
	envconfig.MustProcess("", &config)
	return config
}
