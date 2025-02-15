package implements

import (
	"fmt"
	"sync"
)

type ipMemoryRepositoryImpl struct {
	mu        sync.RWMutex
	hostIpMap map[string][]string
}

func NewIpMemoryRepositoryImpl() *ipMemoryRepositoryImpl {
	return &ipMemoryRepositoryImpl{
		hostIpMap: make(map[string][]string),
	}
}

func (r *ipMemoryRepositoryImpl) FindByHost(host string) ([]string, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	ips, exists := r.hostIpMap[host]
	if exists {
		return ips, nil
	}
	return make([]string, 0), nil
}

func (r *ipMemoryRepositoryImpl) Create(host string, newIps []string) ([]string, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.hostIpMap[host]; exists {
		return nil, fmt.Errorf("host '%s' already exists", host)
	}
	r.hostIpMap[host] = newIps
	return newIps, nil
}

func (r *ipMemoryRepositoryImpl) Update(host string, updateIps []string) ([]string, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.hostIpMap[host]; !exists {
		return nil, fmt.Errorf("host '%s' not found", host)
	}
	r.hostIpMap[host] = updateIps
	return updateIps, nil
}

func (r *ipMemoryRepositoryImpl) Delete(host string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.hostIpMap[host]; !exists {
		return fmt.Errorf("host '%s' not found", host)
	}
	delete(r.hostIpMap, host)
	return nil
}
