package api

/*
  #include <libcgroup.h>
  #cgo LDFLAGS: -lcgroup
*/
import "C"

// Cgroup is the structure describing one or more control groups. The structure
// is opaque to applications.
type Cgroup struct {
	g *C.struct_cgroup
}

// Device structure represents the type storage being provided to the
// data service.
type Device struct {
	// Type could be a string such as "block", "ebs", "nfs" etc.  It is
	// up to the data service to interpret the device type.
	Type string

	// Metadata contains device type specific constraints and information.
	// For example, for an EBS volume type, it can contain the AWS access keys.
	Metadata map[string]string
}

// Node contains details regarding a specific host.  This information
// will be provided to the service being deployed.  This information is
// provided as a file on the host in yaml format.
type Node struct {
	Ips     []string
	Devices []Device

	// Constraints are cgroup restriuctions on the service container.
	Constraints Cgroup

	// ClusterID uniquely identifies the cluster that this
	// data service is part of.
	ClusterID string

	Metadata map[string]string
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
