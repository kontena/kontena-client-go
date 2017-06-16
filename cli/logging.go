package cli

import (
	"os"

	"github.com/op/go-logging"
)

var logFormatter = logging.MustStringFormatter(
	`%{color} %{level:.4s} %{shortpkg}:%{longfunc}:%{color:reset} %{message}`,
)
var logBackend = logging.NewLogBackend(os.Stderr, "", 0)

func init() {
	logging.SetFormatter(logFormatter)
	logging.SetBackend(logBackend)
	logging.SetLevel(logging.DEBUG, "cli")
}

var log = logging.MustGetLogger("cli")

func loggingSetup(options Options) {
	if options.Debug {
		logging.SetLevel(logging.DEBUG, "kontena-cli")
	} else if options.Verbose {
		logging.SetLevel(logging.INFO, "kontena-cli")
	} else if options.Quiet {
		logging.SetLevel(logging.ERROR, "kontena-cli")
	} else {
		logging.SetLevel(logging.WARNING, "kontena-cli")
	}
}

func Logger() *logging.Logger {
	return log
}
