package server

// Server define a server behaviour
type Server interface {
	// Run run the server on the host and port indicated
	Run() error
}
