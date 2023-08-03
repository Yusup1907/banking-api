package manager

import "github.com/Yusup1907/banking-api/src/service"

type ServiceManager interface {
	GetNasabahService() service.NasabahService
}

type serviceManager struct {
	nasabahService service.NasabahService
}

func (sm *serviceManager) GetNasabahService() service.NasabahService {
	return sm.nasabahService
}

func NewServiceManager(repoManager RepositoryManager) ServiceManager {
	nasabahService := service.NewNasabahService(repoManager.GetNasabahRepository())

	return &serviceManager{
		nasabahService: nasabahService,
	}
}
