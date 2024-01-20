package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

// ReadFile reads content from a given file.
func ReadFile(filename string) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Failed to read from file: %s", err)
	}
	fmt.Printf("Contents of %s: \n%s\n", filename, string(data))
}

// WriteLog writes a message to a log file.
func WriteLog(logfile, message string) {
	file, err := os.OpenFile(logfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open log file: %s", err)
	}
	defer file.Close()

	logger := log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Println(message)
}

// MakeHTTPRequest performs a GET request to a URL and logs the response status.
func MakeHTTPRequest(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("HTTP Request failed: %s", err)
	}
	defer resp.Body.Close()

	log.Printf("Received response for %s: %s", url, resp.Status)
}

func main() {
	// Read from a text file
	ReadFile("./data/example.txt")

	// Write to a log file
	WriteLog("./data/activity.log", "Routine task executed.")

	// Perform a scheduled task every 2 seconds
	ticker := time.NewTicker(2 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				// Perform the task
				WriteLog("./data/activity.log", "Scheduled task executed.")
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	// Make an HTTP request
	MakeHTTPRequest("https://httpbin.org/get")

	// The script runs for 1 minute before exiting
	time.Sleep(1 * time.Minute)
	close(quit)
}
