package manager

import (
	"sync"

	"github.com/Yusup1907/banking-api/src/repository"
)

type RepositoryManager interface {
	GetNasabahRepository() repository.NasabahRepository
	GetAccountRepository() repository.AccountRepository
	GetMerchantRepository() repository.MerchantRepository
	GetPaymentRepository() repository.TransactionRepository
}

type repositoryManager struct {
	nasabahRepo  repository.NasabahRepository
	accountRepo  repository.AccountRepository
	merchantRepo repository.MerchantRepository
	paymentRepo  repository.TransactionRepository
}

var onceLoadNasabahRepo sync.Once
var onceLoadAccountRepo sync.Once
var onceLoadMerchantRepo sync.Once
var onceLoadpaymentRepo sync.Once

func (rm *repositoryManager) GetNasabahRepository() repository.NasabahRepository {
	onceLoadNasabahRepo.Do(func() {
		rm.nasabahRepo = repository.NewNasabahRepository()
	})
	return rm.nasabahRepo
}

func (rm *repositoryManager) GetAccountRepository() repository.AccountRepository {
	onceLoadAccountRepo.Do(func() {
		rm.accountRepo = repository.NewAccountRepository()
	})
	return rm.accountRepo
}

func (rm *repositoryManager) GetMerchantRepository() repository.MerchantRepository {
	onceLoadMerchantRepo.Do(func() {
		rm.merchantRepo = repository.NewMerchantRepository()
	})
	return rm.merchantRepo
}

func (rm *repositoryManager) GetPaymentRepository() repository.TransactionRepository {
	onceLoadpaymentRepo.Do(func() {
		rm.paymentRepo = repository.NewTransactionRepository()
	})
	return rm.paymentRepo
}

func NewRepositoryManager() RepositoryManager {
	return &repositoryManager{}
}
