package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type Temperature struct {
	Timestamp   time.Time `json:"timestamp"`
	SensorID    int
	Temperature float64 `json:"temperature"`
}

func main() {
	// Set the mean and standard deviation for the Gaussian distribution
	mean := 27.5
	stdDev := 0.1
	i := 0

	for {
		// Generate a random temperature value based on the Gaussian distribution
		temperature := rand.NormFloat64()*stdDev + mean
		timeStamp := time.Now()
		i++
		// Create a Temperature struct with timestamp and temperature values
		data := Temperature{
			Timestamp:   timeStamp,
			SensorID:    i,
			Temperature: temperature,
		}

		// Create a JSON file
		file, err := os.Create("/tmp/temperature.json")
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}

		// Encode the data struct as JSON and write it to the file
		encoder := json.NewEncoder(file)
		err = encoder.Encode(data)
		if err != nil {
			fmt.Println("Error encoding data:", err)
			return
		}

		file.Close()

		/*fmt.Printf("ID: %d\n", i)
		fmt.Printf("TimeSTamp: %v\n", timeStamp)
		fmt.Printf("Temperature: %.2f\n", temperature)*/

		time.Sleep(1 * time.Second) // Delay for 1 second before generating the next temperature value
	}
}
