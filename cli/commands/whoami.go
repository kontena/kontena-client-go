package commands

import (
	"github.com/kontena/kontena-client-go/cli"
)

type whoami struct {
	URL      string
	Username string
}

type WhoamiCommand struct {
	*cli.CLI
}

// Whoami prints both the URL and username based upon who is currently
// authenticated.
func (cmd WhoamiCommand) Whoami() error {
	var whoami = whoami{
		URL: cmd.Client.String(),
	}

	if user, err := cmd.Client.Users.GetUser(); err != nil {
		return err
	} else {
		whoami.Username = user.Name
	}

	return cli.Print(whoami)
}
