package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

func main() {
	// Create a new cron manager
	c := cron.New()

	// Add a job that runs every minute
	c.AddFunc("* * * * *", func() {
		fmt.Println("Job 1: Running every minute", time.Now())
	})

	// Add a job that runs every day at 9:00 AM
	c.AddFunc("0 9 * * *", func() {
		fmt.Println("Job 2: Running every day at 9:00 AM", time.Now())
	})

	// Start the cron scheduler
	c.Start()

	// Keep the main thread running
	select {}
}
