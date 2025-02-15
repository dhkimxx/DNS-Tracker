package main

import (
	"fmt"
	"sync"
	"time"
	"tracker/config"
	"tracker/tracker"
)

func main() {
	fmt.Println(config.AppConfig.Tracker.TrackingHosts)
	ConcurrencyScheduler()
}

func BasicScheduler() {
	for {
		t := time.Now()
		for _, host := range config.AppConfig.Tracker.TrackingHosts {
			tracker.CheckDNS(host)

		}
		fmt.Println(time.Since(t))
		time.Sleep(1 * time.Second)
	}
}

func ConcurrencyScheduler() {
	for {
		t := time.Now()
		var wg sync.WaitGroup
		for _, host := range config.AppConfig.Tracker.TrackingHosts {
			wg.Add(1)
			go func(host string) {
				defer wg.Done()
				tracker.CheckDNS(host)
			}(host)
		}
		wg.Wait()
		fmt.Println(time.Since(t))
		time.Sleep(1 * time.Second)
	}
}
