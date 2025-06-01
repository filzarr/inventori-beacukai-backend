package main

import (
	"flag"
	"fmt"
	"inventori-beacukai-backend/internal/adapter"
	"inventori-beacukai-backend/internal/infrastructure/config"
	"inventori-beacukai-backend/internal/router"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func main() {
	os.Args = initialize()

	log.Info().Msg("Starting server...")

	cfg := config.Envs

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello")
	})
	adapter.Adapters.Sync(
		adapter.WithRestServer(app),
		adapter.WithPostgres(),
	)

	router.SetupRoutes(app)

	addr := fmt.Sprintf(":%s", cfg.App.Port)
	log.Info().Msgf("Server running at http://localhost%s", addr)

	// Jalankan server dalam goroutine
	go func() {
		if err := app.Listen(addr); err != nil {
			log.Fatal().Err(err).Msg("failed to start server")
		}
	}()

	// Tangkap sinyal shutdown (Ctrl+C, SIGTERM)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info().Msg("Gracefully shutting down server...")

	if err := adapter.Adapters.Unsync(); err != nil {
		log.Error().Err(err).Msg("Error while shutting down adapters")
	} else {
		log.Info().Msg("Server shut down cleanly.")
	}

}

func initialize() (newArgs []string) {
	configPath := flag.String("config_path", "./", "path to config file")
	configFilename := flag.String("config_filename", ".env", "config file name")
	flag.Parse()

	logCfg := *configPath + "/" + *configFilename

	log.Info().Msgf("Initializing configuration with config: %s", logCfg)

	config.Configuration(
		config.WithPath(*configPath),
		config.WithFilename(*configFilename),
	).Initialize()

	adapter.Adapters = &adapter.Adapter{}

	for _, arg := range os.Args {
		if strings.Contains(arg, "config_path") || strings.Contains(arg, "config_filename") {
			continue
		}

		newArgs = append(newArgs, arg)
	}

	return newArgs
}
