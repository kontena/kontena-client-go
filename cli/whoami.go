package cli

type whoami struct {
	URL      string
	Username string
}

func (cli *CLI) Whoami() error {
	var whoami = whoami{
		URL: cli.client.String(),
	}

	if user, err := cli.client.Users.GetUser(); err != nil {
		return err
	} else {
		whoami.Username = user.Name
	}

	return print(whoami)
}
