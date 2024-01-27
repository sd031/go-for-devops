package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "Prints the Hello message",
	Long:  `Prints the Hello message to demonstrate the basic Cobra CLI application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello from my CLI app!")
	},
}

func init() {
	rootCmd.AddCommand(helloCmd)
}
