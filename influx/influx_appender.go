package influx

import (
	"context"
	"fmt"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"

	"vexdev.com/ekz-influx/ekz"
)

// Writes the ekz data to the influx database

type InfluxAppender struct {
	influxClient   influxdb2.Client
	influxWriteAPI api.WriteAPIBlocking
}

func NewInfluxAppender(url string, token string, org string, bucket string) *InfluxAppender {
	client := influxdb2.NewClient(url, token)
	writeAPI := client.WriteAPIBlocking(org, bucket)
	return &InfluxAppender{
		influxClient:   client,
		influxWriteAPI: writeAPI,
	}
}

func (i *InfluxAppender) WriteData(name string, data []ekz.EkzSeriesValues) {
	for _, value := range data {
		// Ekz time is UTC in the format "YYYYMMDDHHmmss"
		parsedTime, err := time.Parse("20060102150405", fmt.Sprint(value.Timestamp))
		if err != nil {
			panic(err)
		}
		p := influxdb2.NewPoint(name,
			map[string]string{
				"unit": "kWh",
			},
			map[string]interface{}{
				"avg": value.Value,
			},
			parsedTime)
		err = i.influxWriteAPI.WritePoint(context.Background(), p)
		if err != nil {
			panic(err)
		}
	}
}
