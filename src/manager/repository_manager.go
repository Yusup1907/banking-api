package manager

import "github.com/Yusup1907/banking-api/src/repository"

type RepositoryManager interface {
	GetNasabahRepository() repository.NasabahRepository
}

type repositoryManager struct {
	nasabahRepo repository.NasabahRepository
}

func (rm *repositoryManager) GetNasabahRepository() repository.NasabahRepository {
	return rm.nasabahRepo
}

func NewRepositoryManager(filePath string) RepositoryManager {
	nasabahRepo := repository.NewNasabahRepository(filePath)

	return &repositoryManager{
		nasabahRepo: nasabahRepo,
	}
}
