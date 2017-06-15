package api

type NodeLabels []string

type Node struct {
	ID            string // Docker/Node ID
	Connected     bool
	CreatedAt     Time
	UpdatedAt     Time
	ConnectedAt   Time
	LastSeenAt    Time
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

type NodePUT struct {
	Labels *NodeLabels `json:"labels,omitempty"`
}