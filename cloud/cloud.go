package cloud

type Addr interface {
	Id() string
	Hostport() string
}

type Clouder interface {
	// Start the image using the given cloud provider to provision the given
	// 'image'. It should return to the address (hostport) of how to access
	// the server
	Start(image string, args map[string]string) (Addr, error)

	// Stop the given server
	Stop(id string) error

	// Returns a list of all the servers currently running.
	List() ([]Addr, error)
}
