package manager

import (
	"sync"

	"github.com/Yusup1907/banking-api/src/service"
)

type ServiceManager interface {
	GetNasabahService() service.NasabahService
	GetAccountService() service.AccountService
	GetMerchantService() service.MerchantService
	GetPaymentService() service.TransactionService
}

type serviceManager struct {
	repositoryManager RepositoryManager
	nasabahService    service.NasabahService
	accountService    service.AccountService
	merchantService   service.MerchantService
	paymentRepo       service.TransactionService
}

var onceLoadNasabahService sync.Once
var onceLoadAccountService sync.Once
var onceLoadMerchantService sync.Once
var onceLoadPaymentService sync.Once

func (sm *serviceManager) GetNasabahService() service.NasabahService {
	onceLoadNasabahService.Do(func() {
		sm.nasabahService = service.NewNasabahService(sm.repositoryManager.GetNasabahRepository())
	})
	return sm.nasabahService
}

func (sm *serviceManager) GetAccountService() service.AccountService {
	onceLoadAccountService.Do(func() {
		sm.accountService = service.NewAccountService(sm.repositoryManager.GetAccountRepository())
	})
	return sm.accountService
}

func (sm *serviceManager) GetMerchantService() service.MerchantService {
	onceLoadMerchantService.Do(func() {
		sm.merchantService = service.NewMerchantService(sm.repositoryManager.GetMerchantRepository())
	})
	return sm.merchantService
}

func (sm *serviceManager) GetPaymentService() service.TransactionService {
	onceLoadPaymentService.Do(func() {
		sm.paymentRepo = service.NewTransactionService(sm.repositoryManager.GetAccountRepository(), sm.repositoryManager.GetPaymentRepository())
	})
	return sm.paymentRepo
}

func NewServiceManager(repositoryManager RepositoryManager) ServiceManager {
	return &serviceManager{
		repositoryManager: repositoryManager,
	}
}
