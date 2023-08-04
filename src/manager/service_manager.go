package manager

import (
	"sync"

	"github.com/Yusup1907/banking-api/src/service"
)

type ServiceManager interface {
	GetNasabahService() service.NasabahService
}

type serviceManager struct {
	repositoryManager RepositoryManager
	nasabahService    service.NasabahService
}

var onceLoadNasabahService sync.Once

func (sm *serviceManager) GetNasabahService() service.NasabahService {
	onceLoadNasabahService.Do(func() {
		sm.nasabahService = service.NewNasabahService(sm.repositoryManager.GetNasabahRepository())
	})
	return sm.nasabahService
}

func NewServiceManager(repositoryManager RepositoryManager) ServiceManager {
	return &serviceManager{
		repositoryManager: repositoryManager,
	}
}
