package api

import (
	"regexp"
)

var ExternalRegistryIDRegexp = regexp.MustCompile(`^([a-z0-9_-]+)/(.+)$`)

type ExternalRegistryID struct {
	Grid string
	Name string
}

func (id *ExternalRegistryID) schema() idSchema {
	return idSchema{&id.Grid, &id.Name}
}

func ParseExternalRegistryID(str string) (id ExternalRegistryID, err error) {
	return id, id.schema().parse(str, ExternalRegistryIDRegexp)
}

func (id *ExternalRegistryID) UnmarshalJSON(buf []byte) error {
	return id.schema().unmarshalJSON(buf)
}

func (id ExternalRegistryID) String() string {
	return id.schema().string()
}

func (id ExternalRegistryID) MarshalJSON() ([]byte, error) {
	return id.schema().marshalJSON()
}
