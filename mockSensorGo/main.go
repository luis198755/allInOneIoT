package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type SensorData struct {
	DeviceID    string  `json:"device_id"`
	ClientID    string  `json:"client_id"`
	SensorType  string  `json:"sensor_type"`
	Temperature float64 `json:"temperature"`
	Timestamp   float64 `json:"timestamp"`
}

func generateRandomTemperature(sensorData *SensorData) {
	// Get the current time
	currentTime := time.Now()
	// Convert the time to seconds since Unix epoch
	seconds := currentTime.Unix()
	// Get the fractional seconds
	fractionalSeconds := float64(currentTime.Nanosecond()) / 1e9
	// Create the timestamp string
	//timestamp := fmt.Sprintf("\"timestamp\": %.6f", float64(seconds)+fractionalSeconds)
	// Create the timestamp string
	timestamp := float64(seconds) + fractionalSeconds

	//fmt.Println(float64(seconds))
	//fmt.Println(seconds)
	//fmt.Println(fractionalSeconds)

	//sensorData.Timestamp = float64(time.Now().Unix())
	sensorData.Timestamp = timestamp
	currentTemp := sensorData.Temperature
	difference := rand.Float64() * 0.5
	addOrSubtract := rand.Intn(2)

	if addOrSubtract == 1 && currentTemp < 35 {
		sensorData.Temperature += difference
	} else if currentTemp > 10 {
		sensorData.Temperature -= difference
	}
}

func main() {
	sensorData := SensorData{
		DeviceID:    "e2e78334",
		ClientID:    "c03d5155",
		SensorType:  "TemperatureGo",
		Temperature: 25,
		Timestamp:   float64(time.Now().Unix()),
	}

	for {
		generateRandomTemperature(&sensorData)

		outputFile, err := os.OpenFile("/tmp/ouput_mock_sensorGo.json", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		defer outputFile.Close()

		jsonData, err := json.Marshal(sensorData)
		if err != nil {
			fmt.Println("Error marshaling JSON:", err)
			return
		}

		_, err = outputFile.WriteString(string(jsonData) + "\n")
		if err != nil {
			fmt.Println("Error writing to file:", err)
		}

		time.Sleep(2 * time.Second)
	}
}
