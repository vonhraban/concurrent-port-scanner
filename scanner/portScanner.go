package scanner

const maxPort = 65535

type PortScanner interface {
	Scan() []int
	Ping(port int) bool
}
