package sdh

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sdh",
	Short: "sdh - system daemon configuration helper",
	Long:  "sdh is an utility to create system configuration files and templates.",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Oh... There was an error while executing command '%s'", err)
		os.Exit(1)
	}
}
