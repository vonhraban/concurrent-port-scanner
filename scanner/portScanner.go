package scanner

const maxPort = 65535

// PortScanner dispatches jobs to ping ports on a given IP
type PortScanner interface {
	Scan() []int
	Ping(port int) bool
}
