package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var printCmd = &cobra.Command{
	Use:   "print [string to print]",
	Short: "Print anything to the screen",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Print: " + strings.Join(args, " "))
	},
}

func init() {
	rootCmd.AddCommand(printCmd)
}
