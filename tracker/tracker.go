package tracker

import (
	"fmt"
	"net"
	"tracker/notifier"
	"tracker/repository"
	"tracker/util"
)

func CheckDNS(host string) {
	repo := repository.GetIpRepository()

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
				err = notifier.Notifyf("New destination ip(%s) of host(%s) is detected\n", lookupIp, host)
				if err != nil {
					fmt.Println(err)
				}
			}
			_, err = repo.Create(host, lookupIps)
			if err != nil {
				panic(err)
			}

		} else {
			isEqual, addedIps, deletedIps := util.IsEqualIpAddress(existingIps, lookupIps)
			if !isEqual {

				for _, addedIp := range addedIps {
					err = notifier.Notifyf("destination ip(%s) of host(%s) is added\n", addedIp, host)
					if err != nil {
						fmt.Println(err)
					}
				}

				for _, deletedIp := range deletedIps {
					err = notifier.Notifyf("destination ip(%s) of host(%s) is deleted\n", deletedIp, host)
					if err != nil {
						fmt.Println(err)
					}
				}

				_, err = repo.Update(host, lookupIps)
				if err != nil {
					panic(err)
				}
			}
		}

	}
}
