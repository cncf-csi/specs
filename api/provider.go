package api

import (
	"net/url"
	"time"
)

// Provider implements a data service provider.  This interface implements the
// union of the the data service's CRUD commands as well as it's
// lifecycle operations.
type Provider interface {
	// Type returns the type of storage the driver provides.
	Type() (string, error)

	// Enumerate all volume that satisfy contraints defined by opts.
	Enumerate(opts map[string]string) ([]*Service, error)

	// ServiceInspect inspects a single volume.
	ServiceInspect(
		volumeID string,
		opts map[string]string,
	) (*Service, error)

	// ServiceCreate creates a new volume.
	ServiceCreate(
		name string,
		opts *ServiceSpec,
	) (*Service, error)

	// ServiceCreateFromSnapshot creates a new volume from an existing snapshot.
	ServiceCreateFromSnapshot(
		snapshotID, volumeName string,
		opts *ServiceSpec,
	) (*Service, error)

	// ServiceCopy copies an existing volume.
	ServiceCopy(
		volumeSrc, volumeDest string,
		opts map[string]string,
	) (*Service, error)

	// ServiceSnapshot snapshots a volume.
	ServiceSnapshot(
		volumeID, snapshotName string,
		opts map[string]string,
	) (*Service, error)

	// ServiceRemove removes a volume.
	ServiceRemove(
		volumeID string,
		opts map[string]string,
	) error

	// ServiceAttach attaches a volume and provides a token clients can use
	// to validate that device has appeared locally.
	ServiceAttach(
		volumeID string,
		opts map[string]string,
	) (*Service, string, error)

	// ServiceDetach detaches a volume.
	ServiceDetach(
		volumeID string,
		opts map[string]string,
	) (*Service, error)

	// Mount mounts volume to specific path
	Mount(
		volumeID, mountpoint string,
		opts map[string]string,
	) error

	// Unmount unmounts volume to specific path
	Unmount(
		volumeID, mountpoint string,
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
