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

func (s *SerialPortScanner) Scan() map[int]bool {
	res := make(map[int]bool)
	for port := 1; port <= 65535; port++ {
		res[port] = s.Ping(port)
	}
	return res
}

func (s *SerialPortScanner) Ping(port int) bool {
	return s.Pinger.Ping(s.IP, port)
}
