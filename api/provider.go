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
	Enumerate(opts map[string]string) ([]*Volume, error)

	// VolumeInspect inspects a single volume.
	VolumeInspect(
		volumeID string,
		opts map[string]string,
	) (*Volume, error)

	// VolumeCreate creates a new volume.
	VolumeCreate(
		name string,
		opts *VolumeSpec,
	) (*Volume, error)

	// VolumeCreateFromSnapshot creates a new volume from an existing snapshot.
	VolumeCreateFromSnapshot(
		snapshotID, volumeName string,
		opts *VolumeSpec,
	) (*Volume, error)

	// VolumeCopy copies an existing volume.
	VolumeCopy(
		volumeSrc, volumeDest string,
		opts map[string]string,
	) (*Volume, error)

	// VolumeSnapshot snapshots a volume.
	VolumeSnapshot(
		volumeID, snapshotName string,
		opts map[string]string,
	) (*Volume, error)

	// VolumeRemove removes a volume.
	VolumeRemove(
		volumeID string,
		opts map[string]string,
	) error

	// VolumeAttach attaches a volume and provides a token clients can use
	// to validate that device has appeared locally.
	VolumeAttach(
		volumeID string,
		opts map[string]string,
	) (*Volume, string, error)

	// VolumeDetach detaches a volume.
	VolumeDetach(
		volumeID string,
		opts map[string]string,
	) (*Volume, error)

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
