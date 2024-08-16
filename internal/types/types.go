package types

type ToDo struct {
	Title    string `json:"title"`
	Note     string `json:"note"`
	Complete bool   `json:"complete"`
}

// Options holds the command line options provided at startup
type Options struct {
	// Path to the configuration file.
	Config string

	// Indicates that the application will run in development mode.
	DevMode bool

	// Logging level. One of:  trace, info, debug, warn, error, fatal, panic
	LogLevel string

	// The REST service port on which to listen
	Port int

	Service string
}
