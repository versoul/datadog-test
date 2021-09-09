package main

import (
	"fmt"
	"github.com/DataDog/datadog-go/statsd"
	"log"
)

func main() {
	fmt.Println("start")

	stats, err := statsd.New("172.17.0.10:8125")
	if err != nil {
		log.Fatal(err)
	}
	defer stats.Close()

	stats.Gauge("custom_campaign.queue_waiting", float64(1212), []string{"tag"}, 1)
	stats.Flush()

	fmt.Println("finish !")
}
