package cmd

import (
	"errors"
	"fmt"
	"os"
	"regexp"

	"github.com/spf13/cobra"
	"github.com/vonhraban/concurrent-port-scanner/scanner"
)

var ip string

var rootCmd = &cobra.Command{
	Use:   "concurrent-port-scanner",
	Short: "Concurrent port scanner",
	Long:  "Go port scanner that uses worker pools",
	Run:   run,
}

func run(cmd *cobra.Command, args []string) {
	// TODO! This validation can move into a custom type
	if err := validateInput(ip); err != nil {
		panic(err)
	}

	pinger := &scanner.NetPinger{}
	res := performScan(&scanner.SerialPortScanner{
		IP:     ip,
		Pinger: pinger,
	})

	for _, port := range res {
		fmt.Printf("%s:%d is open\n", ip, port)
	}
}

func validateInput(ip string) error {
	r := `^(\d{1,3})\.(\d{1,3})\.(\d{1,3})\.(\d{1,3})`
	reg, err := regexp.Compile(r)
	if err != nil {
		return err
	}
	ips := reg.FindStringSubmatch(ip)
	if ips == nil {
		return errors.New("Wrong format of IP address")
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
	rootCmd.Flags().StringVarP(&ip, "ip", "", "", "IP address to scan")
	rootCmd.MarkFlagRequired("ip")
}
