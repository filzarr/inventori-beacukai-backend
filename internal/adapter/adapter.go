package adapter

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

var (
	Adapters *Adapter
)

type Adapter struct {
	RestServer *fiber.App
	Postgres   *sqlx.DB
	Validator  Validator
}

type Option func(adapter *Adapter)
type Validator interface {
	Validate(i any) error
}

func (a *Adapter) Sync(opts ...Option) {
	for _, opt := range opts {
		opt(a)
	}
}

func (a *Adapter) Unsync() error {
	var errs []string

	if a.RestServer != nil {
		if err := a.RestServer.Shutdown(); err != nil {
			errs = append(errs, err.Error())
		}
		log.Info().Msg("Rest server disconnected")
	}

	if a.Postgres != nil {
		if err := a.Postgres.Close(); err != nil {
			errs = append(errs, err.Error())
		}
		log.Info().Msg("Postgres disconnected")
	}

	if len(errs) > 0 {
		err := fmt.Errorf("shutdown errors: %s", strings.Join(errs, "; "))
		log.Error().Err(err).Msg("Adapter shutdown with errors")
		return err
	}

	return nil
}
