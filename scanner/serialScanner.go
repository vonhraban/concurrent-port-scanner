package scanner

type serialPortScanner struct {
	IP     string
	Pinger pinger
}

func NewSerialPortScanner(ip string, pinger pinger) PortScanner {
	return &serialPortScanner{
		IP:     ip,
		Pinger: pinger,
	}
}

func (s *serialPortScanner) Scan() []int {
	res := []int{}
	for port := 1; port <= 65535; port++ {
		if s.Ping(port) {
			res = append(res, port)
		}
	}

	return res
}

func (s *serialPortScanner) Ping(port int) bool {
	return s.Pinger.Ping(s.IP, port)
}
