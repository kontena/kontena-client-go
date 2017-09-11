package cli

import (
	"github.com/fatih/color"
	"github.com/rodaine/table"
)

func init() {
	table.DefaultHeaderFormatter = color.New(color.Underline).SprintfFunc()
	table.DefaultFirstColumnFormatter = color.New(color.Bold).SprintfFunc()
}

// Table returns an object with headers for rows to then be added which is done
// by calling the AddRow func.
func Table(columnHeaders ...interface{}) table.Table {
	return table.New(columnHeaders)
}
