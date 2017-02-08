package api

// Modified version from https://github.com/codedellemc/libstorage as an example.
// This is a minimal definition.
// Ultimately, this will be the simplest and most concise definition that consolidates
// the goodness from muliple volume management drivers.

// Volume definition
type Volume struct {
}

// VolumeSpec are options when creating a new volume.
type VolumeSpec struct {
	AvailabilityZone *string
	IOPS             *int64
	Size             *int64
	Encrypted        *bool
	EncryptionKey    *string
	Options          map[string]string
}

type StorageDriver interface {

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
}
