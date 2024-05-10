package main

import (
	"context"
	"log"
	"testing"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/influxdb"
)

func TestWithRedis(t *testing.T) {
	ctx := context.Background()

	influxdbContainer, err := influxdb.RunContainer(
		ctx, testcontainers.WithImage("influxdb:1.8.10"),
		influxdb.WithDatabase("influx"),
		influxdb.WithUsername("root"),
		influxdb.WithPassword("password"),
	)
	if err != nil {
		log.Fatalf("failed to start container: %s", err)
	}

	// Clean up the container
	defer func() {
		if err := influxdbContainer.Terminate(ctx); err != nil {
			log.Fatalf("failed to terminate container: %s", err)
		}
	}()

}
