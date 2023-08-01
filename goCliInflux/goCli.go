package main

import (
	"context"
	"fmt"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func main() {
	for {
		// Create a new client using an InfluxDB server base URL and an authentication token
		client := influxdb2.NewClient("http://influxdb:8086", "luis-iotdata-auth-token")
		// Use blocking write client for writes to desired bucket
		writeAPI := client.WriteAPIBlocking("tlmx", "iotdata")
		// Create point using full params constructor
		p := influxdb2.NewPoint("stat",
			map[string]string{"unit": "temperature"},
			map[string]interface{}{"avg": 24.5, "max": 45.0},
			time.Now())
		// write point immediately
		writeAPI.WritePoint(context.Background(), p)
		// Create point using fluent style
		p = influxdb2.NewPointWithMeasurement("stat").
			AddTag("unit", "temperature").
			AddField("avg", 23.2).
			AddField("max", 45.0).
			SetTime(time.Now())
		err := writeAPI.WritePoint(context.Background(), p)
		if err != nil {
			panic(err)
		}
		// Or write directly line protocol
		line := fmt.Sprintf("stat,unit=temperature avg=%f,max=%f", 23.5, 45.0)
		err = writeAPI.WriteRecord(context.Background(), line)
		if err != nil {
			panic(err)
		}
		/*
			// Get query client
			queryAPI := client.QueryAPI("tlmx")
			// Get parser flux query result
			result, err := queryAPI.Query(context.Background(), `from(bucket:"iotdata")|> range(start: -1h) |> filter(fn: (r) => r._measurement == "stat")`)
			if err == nil {
				// Use Next() to iterate over query result lines
				for result.Next() {
					// Observe when there is new grouping key producing new table
					if result.TableChanged() {
						fmt.Printf("table: %s\n", result.TableMetadata().String())
					}
					// read result
					fmt.Printf("row: %s\n", result.Record().String())
				}
				if result.Err() != nil {
					fmt.Printf("Query error: %s\n", result.Err().Error())
				}
			} else {
				panic(err)
			}*/
		// Ensures background processes finishes
		client.Close()
		// pauses execution for 3 seconds
		time.Sleep(1 * time.Second)
	}
}
