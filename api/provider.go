package api

type Provider interface {
	// GetServiceType advertises the services offered by this
	// providor on a given node.
	GetServiceType() (DataService, error)

	// GetStat returns the service and network statistics for this providor
	// on a given node.
	GetStat() (ServiceStat, NetStat, error)

	// GetAlerts returns the alerts for this providor on a given node.
	GetAlerts() ([]Alert, error)
}
