package util

import "sort"

func IsEqualIpAddress(existingIps, newIps []string) (isEqual bool, addedIps, deletedIps []string) {
	countMap := make(map[string]int)

	for _, newIp := range ToSortedUniqueStringSlice(newIps) {
		countMap[newIp]++
	}

	for _, exexistingIp := range ToSortedUniqueStringSlice(existingIps) {
		countMap[exexistingIp]--
	}

	for ip, count := range countMap {
		if count > 0 {
			addedIps = append(addedIps, ip)
		} else if count < 0 {
			deletedIps = append(deletedIps, ip)
		}
	}

	isEqual = len(addedIps) == 0 && len(deletedIps) == 0
	return
}

func ToSortedUniqueStringSlice(slice []string) []string {
	uniqueMap := make(map[string]struct{})

	for _, item := range slice {
		uniqueMap[item] = struct{}{}
	}

	uniqueSlice := make([]string, 0)
	for key := range uniqueMap {
		uniqueSlice = append(uniqueSlice, key)
	}

	sort.Strings(uniqueSlice)

	return uniqueSlice
}
