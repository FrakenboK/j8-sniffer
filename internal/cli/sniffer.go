package cli

import (
	"os"

	"github.com/FrakenboK/j8-sniffer/internal/sniffer"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "j8-sniffer",
	Short: "CTF Attack/Defence sniffer by cR4.sh",
	Long:  ``,
	Run:   sniffer.Sniff,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().String("interface", "", "Game interface to listen")
	rootCmd.Flags().String("host", "", "Api host")
	rootCmd.Flags().String("token", "", "Server token")
}
