package cli

import (
	"os"

	"github.com/op/go-logging"
)

func init() {
	logging.SetFormatter(logging.MustStringFormatter(
		`%{color} %{level:.4s} %{module}.%{shortfunc}:%{color:reset} %{message}`,
	))
	logging.SetBackend(logging.NewLogBackend(os.Stderr, "", 0))
}

func makeLogger(options Options, module string) *logging.Logger {
	var logger = logging.MustGetLogger(module)

	if options.Debug {
		logging.SetLevel(logging.DEBUG, module)
	} else if options.Verbose {
		logging.SetLevel(logging.INFO, module)
	} else if options.Quiet {
		logging.SetLevel(logging.ERROR, module)
	} else {
		logging.SetLevel(logging.WARNING, module)
	}

	return logger
}

var log = logging.MustGetLogger("kontena-cli")

func setupLogging(options Options) {
	log = makeLogger(options, "kontena-cli")
}

func Logger() *logging.Logger {
	return log
}
