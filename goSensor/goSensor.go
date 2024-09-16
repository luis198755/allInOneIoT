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
	SensorID    int       `json:"sensorID"`
	Temperature float64   `json:"temperature"`
}

const (
	mean     = 27.5
	stdDev   = 0.1
	filePath = "/tmp/temperature.json"
)

func main() {
	sensorID := 0
	for {
		data := generateTemperatureData(&sensorID)
		err := writeJSONToFile(data)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		time.Sleep(1 * time.Second)
	}
}

func generateTemperatureData(sensorID *int) Temperature {
	temperature := rand.NormFloat64()*stdDev + mean
	*sensorID++
	return Temperature{
		Timestamp:   time.Now(),
		SensorID:    *sensorID,
		Temperature: temperature,
	}
}

func writeJSONToFile(data Temperature) error {
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		return fmt.Errorf("error encoding data: %w", err)
	}

	return nil
}
