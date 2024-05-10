package main

import (
	"fmt"
	"time"

	"vexdev.com/ekz-influx/ekz"
)

func main() {
	ekzReader, err := ekz.NewEkzReader()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = ekzReader.Authenticate("usernameGoesHere", "passwordGoesHere")
	if err != nil {
		fmt.Println(err)
		return
	}

	ekzReader.GetConsumptionData("installationIdGoesHere", time.Now().Add(-24*time.Hour), time.Now())

	// go func() {
	// 	for {
	// 		// Do it
	// 		time.Sleep(15 * time.Minute)
	// 	}
	// }()
}
