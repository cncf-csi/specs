package api

import (
	"net/url"
	"time"
)

// Modified version from https://github.com/codedellemc/libstorage as an example.
// This is a minimal definition.
// Ultimately, this will be the simplest and most concise definition that consolidates
// the goodness from muliple service management drivers.

// Service definition.
type Service struct {
}

// ServiceSpec are options when creating a new data service.
type ServiceSpec struct {
	AvailabilityZone *string
	IOPS             *int64
	Size             *int64
	Encrypted        *bool
	EncryptionKey    *string
	Options          map[string]string
}

type Capability int

const (
	CapabilityEncryption Capability = iota
	CapabilityCompresssion
	CapabilityDeduplication
	CapabilityReplication
	CapabilityDR
	CapabilityMulitAZ
	CapabilityConverged
)

type DataService struct {
	// ServiceType could be a string such as object, block, file.
	ServiceType  string
	Size         uint64
	Iops         uint64
	Capabilities []Capability
}

// Provider implements a data service provider.  This interface implements the
// union of the the data service's CRUD commands as well as it's
// lifecycle operations.
type Provider interface {
	// Type returns the type of storage the driver provides.
	Type() (string, error)

	// Enumerate all service that satisfy contraints defined by opts.
	Enumerate(opts map[string]string) ([]*Service, error)

	// ServiceInspect inspects a single service.
	ServiceInspect(
		ID string,
		opts map[string]string,
	) (*Service, error)

	// ServiceCreate creates a new service.
	ServiceCreate(
		name string,
		opts *ServiceSpec,
	) (*Service, error)

	// ServiceCreateFromSnapshot creates a new service from an existing snapshot.
	ServiceCreateFromSnapshot(
		snapshotID, serviceName string,
		opts *ServiceSpec,
	) (*Service, error)

	// ServiceCopy copies an existing service.
	ServiceCopy(
		serviceSrc, volumeDest string,
		opts map[string]string,
	) (*Service, error)

	// ServiceSnapshot snapshots a service.
	ServiceSnapshot(
		ID, snapshotName string,
		opts map[string]string,
	) (*Service, error)

	// ServiceRemove removes a service.
	ServiceRemove(
		ID string,
		opts map[string]string,
	) error

	// ServiceAttach attaches a service and provides a token clients can use
	// to validate that device has appeared locally.
	ServiceAttach(
		ID string,
		opts map[string]string,
	) (*Service, string, error)

	// ServiceDetach detaches a service.
	ServiceDetach(
		ID string,
		opts map[string]string,
	) (*Service, error)

	// Mount mounts service to specific path
	Mount(
		ID, mountpoint string,
		opts map[string]string,
	) error

	// Unmount unmounts service to specific path
	Unmount(
		ID, mountpoint string,
		opts map[string]string,
	) error

	// GetServiceType advertises the services offered by this
	// providor on a given node.
	GetServiceType() (DataService, error)

	// GetStat returns the service and network statistics for this providor
	// on a given node.
	GetStat() (ServiceStat, NetStat, error)

	// LogStats provides an logging URL for the providor dump
	// service stats to.  An interval of 0 stops the logging.
	LogStats(url url.URL, interval time.Duration) error

	// GetAlerts returns the alerts for this providor on a given node.
	GetAlerts() ([]Alert, error)

	// LogAlerts provides an alerting URL for the providor dump
	// service alerts to.  An interval of 0 stops the logging.
	LogAlerts(url url.URL, interval time.Duration) error
}
