package main

import (
	"fmt"
	"time"
	"tracker/config"
	"tracker/tracker"
)

func main() {
	fmt.Println(config.AppConfig.Tracker.TrackingHosts)
	BasicScheduler()
}

func BasicScheduler() {
	for {
		t := time.Now()
		for _, host := range config.AppConfig.Tracker.TrackingHosts {
			tracker.CheckDNS(host)

		}
		fmt.Printf("[scheduler] spent time: %s\n", time.Since(t))
		time.Sleep(1 * time.Second)
	}
}
