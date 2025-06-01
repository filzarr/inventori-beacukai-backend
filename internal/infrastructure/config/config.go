package config

import (
	"inventori-beacukai-backend/pkg/config"
	"sync"

	"github.com/rs/zerolog/log"
)

var (
	Envs *Config
	once sync.Once
)

type Config struct {
	App struct {
		Name         string `env:"APP_NAME" env-default:"library"`
		Port         string `env:"APP_PORT" env-default:"8080"`
		Environtment string `env:"APP_ENV" env-default:"development"`
	}
	Postgres struct {
		Host     string `env:"POSTGRES_HOST" env-default:"localhost"`
		Port     string `env:"POSTGRES_PORT" env-default:"5432"`
		Username string `env:"POSTGRES_USER" env-default:"postgres"`
		Password string `env:"POSTGRES_PASSWORD" env-default:""`
		Database string `env:"POSTGRES_DB" env-default:"app_dev"`
		SslMode  string `env:"POSTGRES_SSL_MODE" env-default:"require"`
	}
	DB struct {
		ConnectionTimeout int `env:"DB_CONN_TIMEOUT" env-default:"30"`
		MaxOpenCons       int `env:"DB_MAX_OPEN_CONS" env-default:"30"`
		MaxIdleCons       int `env:"DB_MAX_IDLE_CONS" env-default:"10"`
		ConnMaxLifetime   int `env:"DB_CONN_MAX_LIFETIME" env-default:"0"`
	}
	Guard struct {
		JwtPrivateKey string `env:"JWT_PRIVATE_KEY" env-default:"-"`
	}
}

type Option = func(c *Configure) error

type Configure struct {
	path     string
	filename string
}

func Configuration(opts ...Option) *Configure {
	c := &Configure{}

	for _, opt := range opts {
		err := opt(c)
		if err != nil {
			panic(err)
		}
	}
	return c
}

func (c *Configure) Initialize() {
	once.Do(func() {
		Envs = &Config{}
		if err := config.Load(config.Opts{
			Config:    Envs,
			Paths:     []string{c.path},
			Filenames: []string{c.filename},
		}); err != nil {
			log.Fatal().Err(err).Msg("get config error")
		}
	})
}

func WithPath(path string) Option {
	return func(c *Configure) error {
		c.path = path
		return nil
	}
}

func WithFilename(name string) Option {
	return func(c *Configure) error {
		c.filename = name
		return nil
	}
}
