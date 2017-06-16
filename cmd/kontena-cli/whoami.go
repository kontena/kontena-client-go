package main

import (
  "fmt"
	"github.com/spf13/cobra"
)

var whoamiCommand = &cobra.Command{
	Use:   "whoami",
	Short: "Show kontena master user",
	RunE: func(cmd *cobra.Command, args []string) error {
    if user, err := cli.client.Users.GetUser(); err != nil {
      return err
    } else {
      fmt.Printf("URL: %v\n", cli.client.String())
      fmt.Printf("User: %v\n", user.Name)
    }

    return nil
	},
}
