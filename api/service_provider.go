package api

type DataService struct {
	// ServiceType could be a string such as object, block, file.
	ServiceType  string
	Size         uint64
	Iops         uint64
	Capabilities []int
}
