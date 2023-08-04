package manager

import (
	"sync"

	"github.com/Yusup1907/banking-api/src/repository"
)

type RepositoryManager interface {
	GetNasabahRepository() repository.NasabahRepository
}

type repositoryManager struct {
	nasabahRepo repository.NasabahRepository
}

var onceLoadNasabahRepo sync.Once

func (rm *repositoryManager) GetNasabahRepository() repository.NasabahRepository {
	onceLoadNasabahRepo.Do(func() {
		rm.nasabahRepo = repository.NewNasabahRepository()
	})
	return rm.nasabahRepo
}

func NewRepositoryManager() RepositoryManager {
	return &repositoryManager{}
}
