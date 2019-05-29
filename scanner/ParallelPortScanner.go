package scanner

type parallelPortScanner struct {
	IP      string
	Pinger  pinger
	Workers int
}

type scanResult struct {
	port int
	open bool
}

func NewParallelPortScanner(ip string, pinger pinger, workers int) PortScanner {
	return &parallelPortScanner{
		IP:      ip,
		Pinger:  pinger,
		Workers: workers,
	}
}

func (s *parallelPortScanner) worker(id int, jobs <-chan int, results chan<- scanResult) {
	for port := range jobs {
		results <- scanResult{
			port: port,
			open: s.Ping(port),
		}
	}
}

func (s *parallelPortScanner) Scan() []int {
	var res []int
	jobs := make(chan int, max_port)
	results := make(chan scanResult, max_port)

	for w := 0; w <= s.Workers; w++ {
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

func (s *parallelPortScanner) Ping(port int) bool {
	return s.Pinger.Ping(s.IP, port)
}
