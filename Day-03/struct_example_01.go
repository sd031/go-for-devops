package main

import (
	"fmt"
	"math/rand"
	"time"
)

// SensorData represents the data collected by a single sensor
type SensorData struct {
	Temperature float64 // in degrees Celsius
	Humidity    float64 // in percentage
	Pressure    float64 // in hPa
}

func main() {
	// Initialize a pseudo-random number generator
	rand.Seed(time.Now().UnixNano())

	// Define an array to hold data from 3 sensors
	var sensorReadings [3]SensorData

	// Simulate sensor data reading
	for i := range sensorReadings {
		sensorReadings[i] = SensorData{
			Temperature: rand.Float64()*30 + 10,  // Random temperature between 10 and 40
			Humidity:    rand.Float64() * 100,    // Random humidity between 0 and 100%
			Pressure:    rand.Float64()*20 + 980, // Random pressure between 980 and 1000 hPa
		}
	}

	// Process and display the sensor data
	for i, reading := range sensorReadings {
		fmt.Printf("Sensor %d - Temperature: %.2f°C, Humidity: %.2f%%, Pressure: %.2f hPa\n",
			i+1, reading.Temperature, reading.Humidity, reading.Pressure)
	}

	// Additional processing can be done here, such as calculating averages, detecting anomalies, etc.
}
