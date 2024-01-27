package cmd

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"github.com/spf13/cobra"
)

func GetCPUUsage() string {
	percentages, err := cpu.Percent(time.Second, false)
	if err != nil {
		return fmt.Sprintf("Error getting CPU Usage: %s", err)
	}
	return fmt.Sprintf("CPU Usage: %.2f%%", percentages[0])
}

func GetMemoryUsage() string {
	vmStat, err := mem.VirtualMemory()
	if err != nil {
		return fmt.Sprintf("Error getting Memory Usage: %s", err)
	}
	return fmt.Sprintf("Memory Usage: Total: %v, Free: %v, UsedPercent: %.2f%%", vmStat.Total, vmStat.Free, vmStat.UsedPercent)
}

func GetDiskUsage() string {
	diskStat, err := disk.Usage("/")
	if err != nil {
		return fmt.Sprintf("Error getting Disk Usage: %s", err)
	}
	return fmt.Sprintf("Disk Usage: Total: %v, Free: %v, UsedPercent: %.2f%%", diskStat.Total, diskStat.Free, diskStat.UsedPercent)
}

func ExportToFile(result, filePath string) {

	if filePath != "" {
		// Replace newline characters with a space or other delimiter
		formattedData := strings.ReplaceAll(result, "\n", " ")

		timestamp := time.Now().Format("2006-01-02 15:04:05")
		logEntry := fmt.Sprintf("%s - %s\n", timestamp, formattedData)

		// Open the file in append mode, create it if it does not exist
		file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		defer file.Close()

		// Write the log entry to the file
		if _, err := file.WriteString(logEntry); err != nil {
			fmt.Println("Error writing to file:", err)
		} else {
			fmt.Println("Output appended to", filePath)
		}
	}
}

var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "Get CPU, Memory, and Disk usage",
	Run: func(cmd *cobra.Command, args []string) {
		cpuUsage := GetCPUUsage()
		memUsage := GetMemoryUsage()
		diskUsage := GetDiskUsage()

		result := strings.Join([]string{cpuUsage, memUsage, diskUsage}, "\n")
		fmt.Println(result)

		if exportFilePath != "" {
			ExportToFile(result, exportFilePath)
		}
	},
}

func init() {
	rootCmd.AddCommand(statsCmd)
	statsCmd.Flags().StringVarP(&exportFilePath, "export", "e", "", "Export to file (provide file path)")
}
