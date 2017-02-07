package api

// Node contains details regarding a specific host.  This information
// will be provided to the service being deployed.  This information is
// provided as a file on the host in yaml format.
type Node struct {
	Devices   []string
	Ips       []string
	CpuMax    int
	MemMax    int
	ClusterID string
	Metadata  map[string]string
}

// Bootstrap contains information for the scheduler.  It instructs the scheduler
// to deploy a given service on a set of nodes.
type Bootstrap struct {
	ApiVersion      string
	DataServiceName string

	// The Docker image to execute on the hosts.
	Image string

	// Set of nodes to deploy the service on.
	Nodes []Node
}

// Deploy will launch a data service on a set of machines as per the
// bootstrap information.
func Deploy(b Bootstrap) error {
	return nil
}
