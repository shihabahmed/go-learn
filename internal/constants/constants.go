package constants

import "github.com/rs/zerolog"

const ToDoRoute = "/todo"
const ToDoRouteWithParam = "/todo/:id"

const (
	ExitCodeOK int = iota
	ExitCodeServiceStartFailure
	ExitCodeServiceCommandFailure
	ExitCodeInitializeFailure
	ExitCodeIOFailure
	ExitCodeDbConnectFailure
	ExitCodeInvalidFlag
)

const LogDefaultLevel = zerolog.InfoLevel