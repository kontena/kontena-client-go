package api

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
)

type idSchema []*string

func (schema idSchema) parse(str string, re *regexp.Regexp) error {
	if matches := re.FindStringSubmatch(str); matches == nil {
		return fmt.Errorf("Invalid ID: %#v", str)
	} else if len(matches)-1 != len(schema) {
		panic(fmt.Sprintf("Invalid ID regexp=%#v for schema=%#v", re, schema))
	} else {
		for i, part := range schema {
			*part = matches[1+i]
		}
	}

	return nil
}

func (schema idSchema) unmarshalJSON(buf []byte) error {
	var str string

	if err := json.Unmarshal(buf, &str); err != nil {
		return err
	}

	parts := strings.Split(str, "/")

	if len(parts) != len(schema) {
		return fmt.Errorf("Invalid JSON ID: %#v", str)
	}

	for i, part := range schema {
		*part = parts[i]
	}

	return nil
}

func (schema idSchema) string() string {
	var str string

	for i, part := range schema {
		if i != 0 {
			str += "/"
		}
		str += *part
	}

	return str
}

func (schema idSchema) marshalJSON() ([]byte, error) {
	return json.Marshal(schema.string())
}
