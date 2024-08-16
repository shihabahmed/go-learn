package initializers

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/shihabahmed/go-learn/internal/constants"
)

// initLogging sets initial default values for logging
func Logger(fp string) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.TimestampFieldName = "t"

	now := time.Now()
	fp += fmt.Sprintf("/feeds-%d-%02d-%02d.log", now.Year(), now.Month(), now.Day())

	zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
		short := file
		for i := len(file) - 1; i > 0; i-- {
			if file[i] == '/' {
				short = file[i+1:]
				break
			}
		}
		file = short
		return file + ":" + strconv.Itoa(line)
	}

	// for local testing use -d for more readable console output
	if options.DevMode {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	} else {
		logFile, _ := os.OpenFile(
			fp,
			os.O_APPEND|os.O_CREATE|os.O_WRONLY,
			0664,
		)

		multi := zerolog.MultiLevelWriter(os.Stdout, logFile)
		log.Logger = zerolog.New(multi).With().Timestamp().Logger()
	}

	log.Logger = log.With().Caller().Logger()

	// set logging level to command-line flag value, if there is an error use the default
	if logLevel, err := zerolog.ParseLevel(options.LogLevel); err == nil {
		zerolog.SetGlobalLevel(logLevel)
	} else {
		zerolog.SetGlobalLevel(constants.LogDefaultLevel)
		log.Warn().Err(err).Msgf("Unable to set log level to %s", logLevel)
	}

	log.Info().Msgf("Log initialized at %s", fp)
}
