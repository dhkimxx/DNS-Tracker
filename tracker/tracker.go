package tracker

import (
	"fmt"
	"net"
	"strings"
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
			fmt.Println(err)
			return
		}

		if len(existingIps) == 0 {
			err = notifier.Notifyf("[CREATED]\nhost: %s\n\nnew ip list:\n%s", host, strings.Join(lookupIps, "\n"))
			if err != nil {
				fmt.Println(err)
				return
			}
			err = repo.Create(host, lookupIps)
			if err != nil {
				fmt.Println(err)
				return
			}

		} else {
			isEqual, addedIps, deletedIps := util.IsEqualIpAddress(existingIps, lookupIps)
			if !isEqual {
				err = notifier.Notifyf("[UPDATED]\nhost: %s\n\ndeleted ip list:\n%s\n\nadded ip list:\n%s\n\ntotal ip list:\n%s",
					host, strings.Join(deletedIps, "\n"), strings.Join(addedIps, "\n"), strings.Join(lookupIps, "\n"))
				if err != nil {
					fmt.Println(err)
					return
				}

				err = repo.Update(host, lookupIps)
				if err != nil {
					fmt.Println(err)
					return
				}
			}
		}

	}
}
