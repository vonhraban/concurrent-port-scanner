package scanner

const max_port = 65535

type PortScanner interface {
	Scan() []int
	Ping(port int) bool
}
