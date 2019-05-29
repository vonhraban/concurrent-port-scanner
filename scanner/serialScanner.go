package scanner

type SerialPortScanner struct {
	IP     string
	Pinger pinger
}

func NewSerialPortScanner(ip string, pinger pinger) PortScanner {
	return &SerialPortScanner{
		IP:     ip,
		Pinger: pinger,
	}
}

func (s *SerialPortScanner) Scan() []int {
	res := []int{}
	for port := 1; port <= 65535; port++ {
		if s.Ping(port) {
			res = append(res, port)
		}
	}

	return res
}

func (s *SerialPortScanner) Ping(port int) bool {
	return s.Pinger.Ping(s.IP, port)
}
