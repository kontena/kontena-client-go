package main

import (
	"os"

	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("kontena-cli")
var logFormatter = logging.MustStringFormatter(
	`%{color} %{level:.4s} %{shortpkg}:%{longfunc}: %{message}`,
)
var logBackend = logging.NewLogBackend(os.Stderr, "", 0)

func init() {
	logging.SetFormatter(logFormatter)
	logging.SetBackend(logBackend)
	logging.SetLevel(logging.DEBUG, "kontena-cli")
}
