package api

// Modified version from https://github.com/codedellemc/libstorage as an example.
// This is a minimal definition.
// Ultimately, this will be the simplest and most concise definition that consolidates
// the goodness from muliple volume management drivers.

// Service definition
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
