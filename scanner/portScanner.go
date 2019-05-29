package scanner

type PortScanner interface {
	Scan() []int
	Ping(port int) bool
}
