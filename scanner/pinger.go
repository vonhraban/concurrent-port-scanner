package scanner

import (
	"fmt"
	"net"
	"time"
)

type pinger interface {
	Ping(ip string, port int) bool
}

type NetPinger struct{}

func (p *NetPinger) Ping(ip string, port int) bool {
	//fmt.Printf("Scanning %s:%d\n", ip, port)
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", ip, port), 100*time.Millisecond)
	if err != nil {
		return false
	}

	conn.Close()

	return true
}
