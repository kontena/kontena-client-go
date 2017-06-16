package main

import (
	"fmt"
)

type options struct {
	Debug   bool
	Verbose bool
	Quiet   bool
	Token   string
	Grid    string
}

func (options options) requireGrid() (string, error) {
	if options.Grid == "" {
		return "", fmt.Errorf("Requires --grid")
	} else {
		return options.Grid, nil
	}
}
