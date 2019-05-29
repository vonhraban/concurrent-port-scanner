package cmd

import (
	"errors"
	"fmt"
	"os"
	"regexp"

	"github.com/spf13/cobra"
	"github.com/vonhraban/concurrent-port-scanner/scanner"
)

const defaultMode = "parallel"
const defaultWorkers = 5

var ip string
var mode string
var workers int

var rootCmd = &cobra.Command{
	Use:   "concurrent-port-scanner",
	Short: "Concurrent port scanner",
	Long:  "Go port scanner that uses worker pools",
	Run:   run,
}

func run(cmd *cobra.Command, args []string) {
	// TODO! This validation can move into a custom type
	if err := validateInput(); err != nil {
		fmt.Printf("Input error: %s \nUse --help flag for usage \n", err)
		return
	}

	pinger := &scanner.NetPinger{}
	var portScanner scanner.PortScanner

	switch mode {
	case "serial":
		portScanner = &scanner.SerialPortScanner{
			IP:     ip,
			Pinger: pinger,
		}
	case "parallel":
		if workers == 0 {
			workers = defaultWorkers
		}
		portScanner = &scanner.ParallelPortScanner{
			IP:      ip,
			Pinger:  pinger,
			Workers: workers,
		}
	default:
		panic("Unexpected execution mode")
	}

	res := performScan(portScanner)

	for _, port := range res {
		fmt.Printf("%s:%d is open\n", ip, port)
	}
}

func validateInput() error {
	r := `^(\d{1,3})\.(\d{1,3})\.(\d{1,3})\.(\d{1,3})`
	reg, err := regexp.Compile(r)
	if err != nil {
		return err
	}
	ips := reg.FindStringSubmatch(ip)
	if ips == nil {
		return errors.New("Wrong format of IP address")
	}

	if mode != "parallel" && mode != "serial" {
		return errors.New("Mode must be either parallel or serial")
	}

	if mode == "serial" && workers != 0 {
		return errors.New("Number of workers can not be specified for the serial processing")
	}

	return nil
}

func performScan(scanner scanner.PortScanner) []int {
	return scanner.Scan()
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&mode, "mode", "", defaultMode, "Exection type: serial, parallel")
	rootCmd.Flags().IntVarP(&workers, "workers", "", 0, "Number of workers")
	rootCmd.Flags().StringVarP(&ip, "ip", "", "", "IP address to scan")

	rootCmd.MarkFlagRequired("ip")
}
