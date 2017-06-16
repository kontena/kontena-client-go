package main

import (
	"fmt"

	"github.com/urfave/cli"
)

var whoamiCommand = cli.Command{
	Name:  "whoami",
	Usage: "Show kontena master user",
	Action: func(c *cli.Context) error {
		fmt.Printf("URL: %v\n", Client.String())

		if user, err := Client.Users.GetUser(); err != nil {
			return err
		} else {
			fmt.Printf("User: %v\n", user.Name)
		}

		return nil
	},
}
