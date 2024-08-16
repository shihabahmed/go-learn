package initializers

import (
	"os"

	"github.com/rs/zerolog/log"
	"github.com/shihabahmed/go-learn/internal/constants"
	"github.com/shihabahmed/go-learn/internal/types"
	"github.com/shihabahmed/go-learn/internal/utils"
)

func Run() (types.Options, utils.Config) {
	LoadEnvVariables()
	
	// initialize flags
	Flags()

	if err := config.FromFile(options.Config); err != nil {
		log.Fatal().Err(err).Msg("Failed to read configuration file")
		os.Exit(constants.ExitCodeIOFailure)
	}

	Logger(config.LogPath)
	ConnectToDB()

	return options, config
}
