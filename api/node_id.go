package api

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
)

var NodeIDRegexp = regexp.MustCompile(`^([a-z0-9_-]+)/([a-z0-9_-]+)$`)

type NodeID struct {
	Grid string
	Name string
}

func ParseNodeID(str string) (id NodeID, err error) {
	if matches := NodeIDRegexp.FindStringSubmatch(str); matches == nil {
		return id, fmt.Errorf("Invalid NodeID: %#v", str)
	} else {
		id.Grid = matches[1]
		id.Name = matches[2]
	}

	return id, nil
}

func (id *NodeID) UnmarshalJSON(buf []byte) error {
	var str string

	if err := json.Unmarshal(buf, &str); err != nil {
		return err
	}

	parts := strings.Split(str, "/")

	if len(parts) != 2 {
		return fmt.Errorf("Invalid NodeID: %#v", str)
	}

	id.Grid = parts[0]
	id.Name = parts[1]

	return nil
}

func (nodeID NodeID) String() string {
	return fmt.Sprintf("%s/%s", nodeID.Grid, nodeID.Name)
}

func (nodeID NodeID) MarshalJSON() ([]byte, error) {
	return json.Marshal(nodeID.String())
}
