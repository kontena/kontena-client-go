package api

type GridStatsStatsd struct {
	Server string `json:"server"`
	Port   int    `json:"port"`
}

type GridStats struct {
	Statsd *GridStatsStatsd `json:"statsd"`
}

type GridLogsOpts map[string]string

type GridLogs struct {
	Forwarder string       `json:"forwarder"`
	Opts      GridLogsOpts `json:"opts"`
}

type GridTrustedSubnets []string
type GridDefaultAffinity []string

type Grid struct {
	ID              string
	Name            string
	InitialSize     int `json:"initial_size"`
	Token           string
	Stats           GridStats
	DefaultAffinity GridDefaultAffinity `json:"default_affinity,omitempty"`
	TrustedSubnets  GridTrustedSubnets  `json:"trusted_subnets"`
	Subnet          string
	Supernet        string
	Logs            *GridLogs

	NodeCount      int `json:"node_count"`
	ServiceCount   int `json:"service_count"`
	ContainerCount int `json:"container_count"`
	UserCount      int `json:"user_count"`
}

func (grid Grid) String() string {
	return grid.ID
}

type GridsGET struct {
	Grids []Grid
}

type GridPOST struct {
	// Required
	Name        string `json:"name"`
	InitialSize int    `json:"initial_size"`

	// Optional
	Token    string `json:"token,omitempty"`
	Subnet   string `json:"subnet,omitempty"`
	Supernet string `json:"supernet,omitempty"`

	// Optional
	DefaultAffinity *GridDefaultAffinity `json:"default_affinity,omitempty"`
}

type GridPUT struct {
	// Optional
	TrustedSubnets  *GridTrustedSubnets  `json:"trusted_subnets,omitempty"`
	DefaultAffinity *GridDefaultAffinity `json:"default_affinity,omitempty"`
	Stats           *GridStats           `json:"stats,omitempty"`
	Logs            *GridLogs            `json:"logs,omitempty"`
}
