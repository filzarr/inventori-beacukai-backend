package adapter

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

var (
	Adapters *Adapter
)

type Option func(adapter *Adapter)

type Validator interface {
	Validate(i any) error
}

type Adapter struct {
	// Driving Adapters
	RestServer *fiber.App
	WsServer   *http.Server

	//Driven Adapters
	Postgres  *sqlx.DB
	Validator Validator
}

func (a *Adapter) Sync(opts ...Option) {
	for o := range opts {
		opt := opts[o]
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

	if a.WsServer != nil {
		if err := a.WsServer.Close(); err != nil {
			errs = append(errs, err.Error())
		}
		log.Info().Msg("Ws server disconnected")
	}

	if a.Postgres != nil {
		if err := a.Postgres.Close(); err != nil {
			errs = append(errs, err.Error())
		}
		log.Info().Msg("Digihub Postgres disconnected")
	}

	// if a.VenamonGolog != nil {
	// 	a.VenamonGolog.Stop()
	// 	log.Info().Msg("Venamon Golog disconnected")
	// }

	if len(errs) > 0 {
		err := fmt.Errorf("%s", strings.Join(errs, "\n"))
		log.Error().Msgf("Error while disconnecting adapters: %v", err)
		return err
	}

	return nil
}
