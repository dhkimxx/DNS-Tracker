package repository

type ipRepository interface {
	FindByHost(host string) ([]string, error)
	Create(host string, newIps []string) ([]string, error)
	Update(host string, updateIps []string) ([]string, error)
	Delete(host string) error
}

func NewIpRepository() ipRepository {
	return newIpMemoryRepositoryImpl()
}
