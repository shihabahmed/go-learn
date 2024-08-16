package initializers

import (
	"flag"
	"fmt"
	"os"

	"github.com/shihabahmed/go-learn/internal/constants"
)

func Flags() {
	initFlags()
	validateFlags()
}

// initFlags defines and parses out the command line flags
func initFlags() {
	flag.StringVar(&options.Config, "config", "./config.yaml", "Path to the configuration file")
	flag.BoolVar(&options.DevMode, "devMode", false, "Outputs logs to console in a more human friendly format")
	var help = flag.Bool("help", false, "Shows this usage screen")
	flag.StringVar(&options.LogLevel, "logLevel", "info", "Logging level. One of:  trace, info, debug, warn, error, fatal, panic")
	flag.IntVar(&options.Port, "port", 5000, "The REST service port on which to listen")
	flag.Usage = usage
	flag.Parse()

	if *help {
		usage()
		os.Exit(constants.ExitCodeOK)
	}
}

// validateFlags enforces the use of mandatory flags and validates their values
func validateFlags() error {
	if options.Config == "" {
		return fmt.Errorf("no config file path was specified. A config file is mandatory")
	}

	return nil
}

// usage displays the help information for the commandline application
func usage() {
	fmt.Fprintf(os.Stderr, "Change me v%s, build %s\n", Version, Build)
	fmt.Fprintln(os.Stderr, "Usage: changeme [options...]")
	flag.PrintDefaults()
}