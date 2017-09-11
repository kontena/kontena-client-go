package cli

import (
	"os"

	"gopkg.in/yaml.v2"
)

// Print marshals the given object into YAML and then prints it to stdout.
func Print(object interface{}) error {
	if out, err := yaml.Marshal(object); err != nil {
		return err
	} else {
		os.Stdout.Write(out)
	}

	return nil
}
