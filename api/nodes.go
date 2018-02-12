package api

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"time"
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

type NodeLabels []string

// Node represents a Kontena node, (a virtual or physical machine) which will
// run the Kontena Agent where services can be deployed.
type Node struct {
	ID            NodeID // grid/name
	NodeID        string `json:"node_id"` // Docker ID
	Connected     bool
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	ConnectedAt   time.Time `json:"connected_at"`
	LastSeenAt    time.Time `json:"last_seen_at"`
	Name          string
	OS            string
	EngineRootDir string `json:"engine_root_dir"`
	Driver        string
	// network_drivers
	// volume_drivers
	KernelVersion string `json:"kernel_version"`
	Labels        NodeLabels
	MemTotal      int `json:"mem_total"`
	MemLimit      int `json:"mem_limit"`
	CPUs          int
	PublicIP      string   `json:"public_ip"`
	PrivateIP     string   `json:"private_ip"`
	OverlayIP     string   `json:"overlay_ip"`
	AgentVersion  string   `json:"agent_version"`
	DockerVersion string   `json:"docker_version"`
	PeerIPs       []string `json:"peer_ips"`
	NodeNumber    int      `json:"node_number"`
	InitialMember bool     `json:"initial_member"`
	Grid          Grid
	// resource_usage
}

type Nodes []Node

type NodesGET struct {
	Nodes Nodes
}

type NodePOST struct {
	Name   string      `json:"name"`
	Token  string      `json:"token,omitempty"`
	Labels *NodeLabels `json:"labels,omitempty"`
}

type NodePUT struct {
	Labels *NodeLabels `json:"labels,omitempty"`
}

type NodeToken struct {
	ID    NodeID `json:"id"`
	Token string `json:"token"`
}

type NodeTokenParams struct {
	ResetConnection bool `json:"reset_connection"`
}

type NodeTokenPUT struct {
	Token string `json:"token,omitempty"`
	NodeTokenParams
}

type NodeTokenDELETE struct {
	NodeTokenParams
}
