package main

import (
	"fmt"
	"os"
	"time"

	"vexdev.com/ekz-influx/ekz"
	"vexdev.com/ekz-influx/influx"
)

func main() {
	username := os.Getenv("EKZ_USERNAME")
	password := os.Getenv("EKZ_PASSWORD")
	installationId := os.Getenv("EKZ_INSTALLATION_ID")
	influxUrl := os.Getenv("INFLUX_URL")
	influxToken := os.Getenv("INFLUX_TOKEN")
	influxOrg := os.Getenv("INFLUX_ORG")
	influxBucket := os.Getenv("INFLUX_BUCKET")

	// Check if all environment variables are set
	if username == "" || password == "" || installationId == "" || influxUrl == "" || influxToken == "" || influxOrg == "" || influxBucket == "" {
		fmt.Println("Please set the following environment variables:")
		fmt.Println("EKZ_USERNAME")
		fmt.Println("EKZ_PASSWORD")
		fmt.Println("EKZ_INSTALLATION_ID")
		fmt.Println("INFLUX_URL")
		fmt.Println("INFLUX_TOKEN")
		fmt.Println("INFLUX_ORG")
		fmt.Println("INFLUX_BUCKET")
		return
	}

	ekzReader, err := ekz.NewEkzReader()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = ekzReader.Authenticate(username, password)
	if err != nil {
		fmt.Println(err)
		return
	}

	influxAppender := influx.NewInfluxAppender(influxUrl, influxToken, influxOrg, influxBucket)
	lastFetchedTimestamp := 0

	for {
		data, err := ekzReader.GetConsumptionData(installationId, time.Now().Add(-24*time.Hour), time.Now())
		if err != nil {
			// Retry once
			err = ekzReader.Authenticate(username, password)
			if err != nil {
				panic(err)
			}
			data, err = ekzReader.GetConsumptionData(installationId, time.Now().Add(-24*time.Hour), time.Now())
			if err != nil {
				panic(err)
			}
		}

		allValues := data.GetAllValidValues()
		maxTimestamp := getMaxTimestamp(allValues)
		if len(allValues) == 0 || maxTimestamp <= lastFetchedTimestamp {
			println("No new values found")
			time.Sleep(15 * time.Minute)
			continue
		}
		lastFetchedTimestamp = maxTimestamp
		influxAppender.WriteData("power", allValues)
		println("DataPoints written to influx power: ", len(allValues))
		htValues := data.GetValidHtValues()
		influxAppender.WriteData("power_ht", htValues)
		println("DataPoints written to influx power_ht: ", len(htValues))
		ntValues := data.GetValidNtValues()
		influxAppender.WriteData("power_nt", ntValues)
		println("DataPoints written to influx power_nt: ", len(ntValues))
		time.Sleep(15 * time.Minute)
	}
}

func getMaxTimestamp(values []ekz.EkzSeriesValues) int {
	max := 0
	for _, value := range values {
		if value.Timestamp > max {
			max = value.Timestamp
		}
	}
	return max
}
