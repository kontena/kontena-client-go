package test_data

import "github.com/kontena/terraform-provider-kontena/api"

var Node = api.Node{
	ID:            "VAHL:WEKY:EVK7:7WUJ:2SQY:B5T7:36TT:NTUL:7CHW:XJA5:KBHL:RSTU",
	Connected:     true,
	Name:          "terraform-test-node1",
	OS:            "Container Linux by CoreOS 1353.8.0 (Ladybug)",
	EngineRootDir: "/var/lib/docker",
	Driver:        "overlay",
	KernelVersion: "4.9.24-coreos",
	Labels: api.NodeLabels{
		"provider=digitalocean",
		"node-index=1",
		"region=fra1",
		"az=fra1",
	},
	MemTotal:      1045307392,
	MemLimit:      0,
	CPUs:          1,
	PublicIP:      "138.68.79.119",
	PrivateIP:     "138.68.79.119",
	OverlayIP:     "10.81.0.1",
	AgentVersion:  "1.3.0",
	DockerVersion: "1.12.6",
	PeerIPs:       []string{"138.68.94.29"},
	NodeNumber:    1,
	InitialMember: true,
	Grid: api.Grid{
		ID:          "test",
		Name:        "test",
		InitialSize: 1,
		Token:       "baCPXomhZucwpRVmNcdPBTvQ9YinRXZyxODrK3UBMxd5Rjjuw0rx5MdG6YzN9q/AwpU1vch0k8rcaG/4qt+mpw==",
		Stats: api.GridStats{
			Statsd: nil,
		},
		DefaultAffinity: nil,
		TrustedSubnets:  api.GridTrustedSubnets{},
		Subnet:          "10.81.0.0/16",
		Supernet:        "10.80.0.0/12",
		Logs:            nil,
	},
}
