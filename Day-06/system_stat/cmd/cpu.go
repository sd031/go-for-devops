package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	// Replace with your module path
)

var cpuCmd = &cobra.Command{
	Use:   "cpu",
	Short: "Get CPU usage",
	Run: func(cmd *cobra.Command, args []string) {
		result := GetCPUUsage()
		fmt.Println(result)
		ExportToFile(result, exportFilePath)
	},
}

func init() {
	rootCmd.AddCommand(cpuCmd)
	cpuCmd.Flags().StringVarP(&exportFilePath, "export", "e", "", "Export to file (provide file path)")
}
