package scanner

type ChannelPortScanner struct {
	IP     string
	Pinger pinger
}

type scanResult struct {
	port int
	open bool
}

func NewChannelPortScanner(ip string, pinger pinger) PortScanner {
	return &SerialPortScanner{
		IP:     ip,
		Pinger: pinger,
	}
}

func (s *ChannelPortScanner) worker(id int, jobs <-chan int, results chan<- scanResult) {
	for port := range jobs {
		results <- scanResult{
			port: port,
			open: s.Ping(port),
		}
	}
}

func (s *ChannelPortScanner) Scan() []int {
	var res []int
	jobs := make(chan int, max_port)
	results := make(chan scanResult, max_port)

	for w := 0; w <= 500; w++ {
		go s.worker(w, jobs, results)
	}

	for i := 1; i < max_port; i++ {
		jobs <- i
	}

	for i := 1; i < max_port; i++ {
		scanResult := <-results
		if scanResult.open {
			res = append(res, scanResult.port)
		}
	}

	return res
}

func (s *ChannelPortScanner) Ping(port int) bool {
	return s.Pinger.Ping(s.IP, port)
}
