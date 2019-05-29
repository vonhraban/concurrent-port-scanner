# concurrent-port-scanner

A port scanner written in Go. It supports serial and parallel
processing with a configurable number of workers in the pool.

### Running
```
$ ./concurrent-port-scanner --help
Go port scanner that uses worker pools

Usage:
  concurrent-port-scanner [flags]

Flags:
  -h, --help          help for concurrent-port-scanner
      --ip string     IP address to scan
      --mode string   Exection type: serial, parallel (default "parallel")
      --workers int   Number of workers
```

### Benchmarking
Performed on 
- MacOS 10.14.2
- 2.9 GHz Intel Core i7 4 cores (8 logical)
- 16GB 2133 MHz LPDDR3

#### Serial processing
```
$ time ./concurrent-port-scanner --ip 127.0.0.1 --mode serial >> /dev/null

real    8m5.107s
user    0m3.969s
sys     0m5.892s
```
#### Parallel processing
```
$ time ./concurrent-port-scanner --ip 127.0.0.1 --mode parallel --workers 5>> /dev/null

real	5m57.152s
user	0m3.830s
sys	0m6.184s

$ time ./concurrent-port-scanner --ip 127.0.0.1 --mode parallel --workers 1000 >> /dev/null

real	0m7.154s
user	0m2.706s
sys	0m6.195s

$ time ./concurrent-port-scanner --ip 127.0.0.1 --mode parallel --workers 10000 >> /dev/null

real	0m2.636s
user	0m2.122s
sys	0m3.004s

$ time ./concurrent-port-scanner --ip 127.0.0.1 --mode parallel --workers 100000 >> /dev/null

real	0m5.932s
user	0m4.751s
sys	0m5.624s
```

I plan to use go testing framework benchmarking functionality to build more data and plot a graph.

### Testing
The usual `go test ./...`
