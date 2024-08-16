package initializers

import (
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func LoadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal().Msg("Error loading .env file.");
	} else {
		log.Info().Msg("Successfully loaded .env file.")
	}
}