package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	// Replace with your module path
)

var diskCmd = &cobra.Command{
	Use:   "disk",
	Short: "Get Disk usage",
	Run: func(cmd *cobra.Command, args []string) {
		result := GetDiskUsage()
		fmt.Println(result)
		ExportToFile(result, exportFilePath)
	},
}

func init() {
	rootCmd.AddCommand(diskCmd)
	diskCmd.Flags().StringVarP(&exportFilePath, "export", "e", "", "Export to file (provide file path)")
}
