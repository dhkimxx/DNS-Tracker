package repository

import implements "tracker/repository/impl"

type IpRepository interface {
	FindByHost(host string) ([]string, error)
	Create(host string, newIps []string) error
	Update(host string, updateIps []string) error
	Delete(host string) error
}

func GetIpRepository() IpRepository {
	return implements.GetRedisIpRepository("localhost:6379", "", 0)
}
