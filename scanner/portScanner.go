package scanner

type PortScanner interface {
	Scan() map[int]bool
	Ping(port int) bool
}
