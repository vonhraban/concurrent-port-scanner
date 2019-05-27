package scanner

import (
	"fmt"
	"net"
	"time"
)

type scanner interface {
	Scan(ip string, port string) map[int]bool
}

type PortScanner struct {
	IP string
}

func (s *PortScanner) Scan() map[int]bool {
	res := make(map[int]bool)
	for port := 1; port <= 65535; port++ {
		res[port] = s.ping(port)
	}
	return res
}

func (s *PortScanner) ping(port int) bool {
	fmt.Printf("Scanning %s:%d\n", s.IP, port)
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", s.IP, port), 100 * time.Millisecond)
	if err != nil {
		return false
	}

 	conn.Close()

	return true
}
