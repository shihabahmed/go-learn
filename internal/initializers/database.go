package initializers

import (
	"os"

	"github.com/rs/zerolog/log"
	"github.com/shihabahmed/go-learn/internal/constants"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error

	if DB, err = config.DbConfig.Connect(); err != nil {
		log.Error().AnErr("error", err).Msgf("Failed to open db at %s", options.Config)
		os.Exit(constants.ExitCodeDbConnectFailure)
	}
}
