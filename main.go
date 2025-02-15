package main

import (
	"fmt"
	"net"
	"time"
	"tracker/config"
	"tracker/repository"
	"tracker/util"
)

func main() {

	fmt.Println(config.AppConfig.Tracker.TrackingHosts)
	repo := repository.NewIpRepository()

	for {
		for _, host := range config.AppConfig.Tracker.TrackingHosts {
			t := time.Now()
			ips, err := net.LookupIP(host)
			if err != nil {
				panic(err)
			}

			if len(ips) > 0 {

				lookupIps := make([]string, len(ips))
				for i, ip := range ips {
					lookupIps[i] = ip.String()
				}

				existingIps, err := repo.FindByHost(host)
				if err != nil {
					panic(err)
				}

				if len(existingIps) == 0 {
					for _, lookupIp := range lookupIps {
						fmt.Printf("New destination ip(%s) of host(%s) is detected\n", lookupIp, host)
					}
					_, err = repo.Create(host, lookupIps)
					if err != nil {
						panic(err)
					}

				} else {
					// compare ips
					isEqual, addedIps, deletedIps := util.IsEqualIpAddress(existingIps, lookupIps)
					if !isEqual {

						for _, addedIp := range addedIps {
							fmt.Printf("destination ip(%s) of host(%s) is added\n", addedIp, host)
						}

						for _, deletedIp := range deletedIps {
							fmt.Printf("destination ip(%s) of host(%s) is deleted\n", deletedIp, host)
						}

						_, err = repo.Update(host, lookupIps)
						if err != nil {
							panic(err)
						}
					}
				}

				fmt.Println(time.Since(t))
			}

		}
		time.Sleep(1 * time.Second)
	}
}
