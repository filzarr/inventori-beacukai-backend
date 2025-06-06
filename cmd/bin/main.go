package main

import (
	"flag"
	"inventori-beacukai-backend/cmd"
	"inventori-beacukai-backend/internal/adapter"
	"inventori-beacukai-backend/internal/infrastructure/config"
	"os"
	"strings"

	"github.com/rs/zerolog/log"
)

func main() {
	os.Args = initialize()

	serverCmd := flag.NewFlagSet("server", flag.ExitOnError)
	seedCmd := flag.NewFlagSet("seed", flag.ExitOnError)

	if len(os.Args) < 2 {
		log.Info().Msg("No command provided, defaulting to 'server'")
		cmd.RunServer(serverCmd, os.Args[1:])
		return
	}

	switch os.Args[1] {
	case "seed":
		cmd.RunSeed(seedCmd, os.Args[2:])
	case "server":
		cmd.RunServer(serverCmd, os.Args[2:])
	default:
		log.Info().Msg("Invalid command provided, defaulting to 'server' with provided flags")
		if strings.HasPrefix(os.Args[1], "-") {
			cmd.RunServer(serverCmd, os.Args[1:])
		} else {
			cmd.RunServer(serverCmd, os.Args[2:])
		}
	}
}

func initialize() (newArgs []string) {
	configPath := flag.String("config_path", "./", "path to config file")
	configFilename := flag.String("config_filename", ".env", "config file name")
	flag.Parse()

	configFullPath := *configPath
	if !strings.HasSuffix(configFullPath, "/") {
		configFullPath += "/"
	}
	configFullPath += *configFilename

	log.Info().Msgf("Initializing configuration with config: %s", configFullPath)

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
