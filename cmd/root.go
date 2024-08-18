package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	Timeout   int64
	StartPort int
	EndPort   int
	Host      string
)

func init() {
	rootCmd.Flags().Int64VarP(&Timeout, "timeout", "t", 200, "timeout to echo the input")
	rootCmd.Flags().IntVarP(&StartPort, "start", "s", 0, "start port to echo the input")
	rootCmd.Flags().IntVarP(&EndPort, "end", "e", 65535, "end port to echo the input")
	rootCmd.Flags().StringVarP(&Host, "url", "u", "127.0.0.1", "host to echo the input")
}

var rootCmd = &cobra.Command{
	Use: "pst",
	Run: func(cmd *cobra.Command, args []string) {
		Scan()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
