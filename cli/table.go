package cli

import (
	"github.com/fatih/color"
	"github.com/rodaine/table"
)

func init() {
	table.DefaultHeaderFormatter = color.New(color.Underline).SprintfFunc()
	table.DefaultFirstColumnFormatter = color.New(color.Bold).SprintfFunc()
}

func makeTable(columnHeaders ...interface{}) table.Table {
	return table.New(columnHeaders)
}
