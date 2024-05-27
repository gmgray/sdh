package sdh

import (
	"fmt"

	"github.com/gmgray/sdh/pkg/sdh"
	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:     "test",
	Aliases: []string{"tst"},
	Short:   "Tests input string",
	Long:    "Tests input string, copying it to output",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		output := sdh.Test(args[0])
		fmt.Println(output)
	},
}

func init() {
	rootCmd.AddCommand(testCmd)
}
