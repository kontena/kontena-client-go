package cli

import (
	"os"

	"gopkg.in/yaml.v2"
)

func print(object interface{}) error {
	if out, err := yaml.Marshal(object); err != nil {
		return err
	} else {
		os.Stdout.Write(out)
	}

	return nil
}
