package api

import (
	"regexp"
)

var NodeIDRegexp = regexp.MustCompile(`^([a-z0-9_-]+)/([a-z0-9_-]+)$`)

type NodeID struct {
	Grid string
	Name string
}

func (id *NodeID) schema() idSchema {
	return idSchema{&id.Grid, &id.Name}
}

func ParseNodeID(str string) (id NodeID, err error) {
	return id, id.schema().parse(str, NodeIDRegexp)
}

func (id *NodeID) UnmarshalJSON(buf []byte) error {
	return id.schema().unmarshalJSON(buf)
}

func (id NodeID) String() string {
	return id.schema().string()
}

func (id NodeID) MarshalJSON() ([]byte, error) {
	return id.schema().marshalJSON()
}
