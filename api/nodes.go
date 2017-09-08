package api

import (
	"time"
)

type NodeLabels []string

// Node represents a Kontena node, (a virtual or physical machine) which will
// run the Kontena Agent where services can be deployed.
type Node struct {
	ID            string // grid/name
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
	ID    string `json:"id"`
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
