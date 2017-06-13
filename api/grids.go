package api

type GridStatsStatsd struct {
	Server string `json:"server"`
	Port   string `json:"port"` // XXX: really?
}

type GridStats struct {
	Statsd *GridStatsStatsd `json:"statsd"`
}

type GridLogs struct {
	Forwarder string            `json:"forwarder"`
	Opts      map[string]string `json:"opts"`
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
}

func (grid Grid) String() string {
	return grid.ID
}

type GridPOST struct {
	// Required
	Name        string `json:"name"`
	InitialSize int    `json:"initial_size"`

	// Optional
	Token           *string              `json:"token,omitempty"`
	DefaultAffinity *GridDefaultAffinity `json:"default_affinity,omitempty"`
	Subnet          *string              `json:"subnet,omitempty"`
	Supernet        *string              `json:"supernet,omitempty"`
}

type GridPUT struct {
	TrustedSubnets  *GridTrustedSubnets  `json:"trusted_subnets,omitempty"`
	DefaultAffinity *GridDefaultAffinity `json:"default_affinity,omitempty"`
	Stats           *GridStats           `json:"stats,omitempty"`
	Logs            *GridLogs            `json:"logs,omitempty"`
}
