package scanner

type serialPortScanner struct {
	IP     string
	Pinger pinger
}

// NewSerialPortScanner creates a new instance of serial port scanner
func NewSerialPortScanner(ip string, pinger pinger) PortScanner {
	return &serialPortScanner{
		IP:     ip,
		Pinger: pinger,
	}
}

func (s *serialPortScanner) Scan() []int {
	res := []int{}
	for port := 1; port <= maxPort; port++ {
		if s.Ping(port) {
			res = append(res, port)
		}
	}

	return res
}

func (s *serialPortScanner) Ping(port int) bool {
	return s.Pinger.Ping(s.IP, port)
}
