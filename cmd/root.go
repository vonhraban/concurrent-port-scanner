package cmd

import (
	"fmt"
	"os"

  "github.com/spf13/cobra"
)

var ip string


var rootCmd = &cobra.Command{ 
	Use:   "concurrent-port-scanner",
	Short: "Concurrent port scanner",
	Long: "Go port scanner that uses worker pools",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%+v", ip)
	},
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
