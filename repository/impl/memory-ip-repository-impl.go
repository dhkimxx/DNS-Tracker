package implements

import (
	"fmt"
	"sync"
)

type IpMemoryRepositoryImpl struct {
	mu        sync.RWMutex
	hostIpMap map[string][]string
}

var (
	memoryIpRepoInstance *IpMemoryRepositoryImpl
	memoryIpRepoOnce     sync.Once
)

func GetIpMemoryRepositoryImpl() *IpMemoryRepositoryImpl {
	memoryIpRepoOnce.Do(func() {
		memoryIpRepoInstance = &IpMemoryRepositoryImpl{
			hostIpMap: make(map[string][]string),
		}
	})
	return memoryIpRepoInstance
}

func (r *IpMemoryRepositoryImpl) FindByHost(host string) ([]string, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	ips, exists := r.hostIpMap[host]
	if exists {
		return ips, nil
	}
	return make([]string, 0), nil
}

func (r *IpMemoryRepositoryImpl) Create(host string, newIps []string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.hostIpMap[host]; exists {
		return fmt.Errorf("host '%s' already exists", host)
	}
	r.hostIpMap[host] = newIps
	return nil
}

func (r *IpMemoryRepositoryImpl) Update(host string, updateIps []string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.hostIpMap[host]; !exists {
		return fmt.Errorf("host '%s' not found", host)
	}
	return nil
}

func (r *IpMemoryRepositoryImpl) Delete(host string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.hostIpMap[host]; !exists {
		return fmt.Errorf("host '%s' not found", host)
	}
	delete(r.hostIpMap, host)
	return nil
}
