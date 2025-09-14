package cmd

import (
	"flag"
	"inventori-beacukai-backend/internal/adapter"
	"inventori-beacukai-backend/internal/infrastructure"
	"inventori-beacukai-backend/internal/infrastructure/config"
	"inventori-beacukai-backend/internal/middleware"
	"inventori-beacukai-backend/internal/route"
	"inventori-beacukai-backend/pkg/validator"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func RunServer(cmd *flag.FlagSet, args []string) {
	var (
		envs        = config.Envs
		flagAppPort = cmd.String("port", "3000", "Application port")
		SERVER_PORT string
	)

	logLevel, err := zerolog.ParseLevel(envs.App.LogLevel)
	if err != nil {
		logLevel = zerolog.InfoLevel
	}

	if err := cmd.Parse(args); err != nil {
		log.Fatal().Err(err).Msg("Error while parsing flags")
	}

	if envs.App.Port != "" {
		SERVER_PORT = envs.App.Port
	} else {
		SERVER_PORT = *flagAppPort
	}

	app := fiber.New()

	// Application Local Storage
	err = os.MkdirAll(envs.App.LocalStoragePublicPath, os.ModePerm)
	if err != nil {
		log.Fatal().Err(err).Msg("Error while creating local storage directory")
	}

	// Application private storage
	err = os.MkdirAll(envs.App.LocalStoragePrivatePath, 0700)
	if err != nil {
		log.Fatal().Err(err).Msg("Error while creating local storage directory")
	}

	// Application Middlewares
	if envs.App.Environtment == "production" {
		app.Use(limiter.New(limiter.Config{
			Max:        50,
			Expiration: 30 * time.Second,
		}))
	}
	allowedOrigins := "*"
	if envs.App.Environtment != "production" {
		allowedOrigins = config.Envs.FrontendURL.ClientBaseURL
	}
	app.Use(cors.New(cors.Config{
		AllowOrigins:     allowedOrigins,
		AllowCredentials: true,
		AllowMethods:     "GET,POST,PUT,DELETE,PATCH,OPTIONS,HEAD",
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin,Authorization",
	}))

	// Access log middleware
	app.Use(middleware.AccessLog)
	// End Application Middlewares

	adapter.Adapters.Sync(
		adapter.WithRestServer(app),
		adapter.WithPostgres(),
		adapter.WithValidator(validator.NewValidator()),
	)

	infrastructure.InitializeLogger(envs.App.Environtment, envs.App.LogFile, logLevel)
	app.Static("api/storage/public", envs.App.LocalStoragePublicPath)
	metricTitle := envs.App.Name + " " + envs.App.Environtment + " " + "Metrics"
	app.Get("/metrics", monitor.New(monitor.Config{Title: metricTitle}))
	route.SetupRoutes(app)

	// Run server in goroutine
	go func() {
		log.Info().Msgf("Server is running on port %s", SERVER_PORT)
		if err := app.Listen(":" + SERVER_PORT); err != nil {
			log.Fatal().Msgf("Error while starting server: %v", err)
		}
	}()
	// End Run server in goroutine

	// Handle graceful shutdown
	quit := make(chan os.Signal, 1)

	shutdownSignals := []os.Signal{os.Interrupt, syscall.SIGTERM, syscall.SIGINT}
	if runtime.GOOS == "windows" {
		shutdownSignals = []os.Signal{os.Interrupt}
	}

	signal.Notify(quit, shutdownSignals...)
	<-quit
	log.Info().Msg("Server is shutting down ...")

	err = adapter.Adapters.Unsync()
	if err != nil {
		log.Error().Msgf("Error while closing adapters: %v", err)
	}

	log.Info().Msg("Server gracefully stopped")
}
