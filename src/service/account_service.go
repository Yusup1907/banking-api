package service

import (
	"github.com/Yusup1907/banking-api/src/model"
	"github.com/Yusup1907/banking-api/src/repository"
)

type AccountService interface {
	AddAccount(account *model.Account) error
}

type accountService struct {
	accountRepo repository.AccountRepository
}

func (s *accountService) AddAccount(account *model.Account) error {
	err := s.accountRepo.AddAccount(account)
	if err != nil {
		// Handle any error from the repository, if needed
		return err
	}
	return nil
}

func NewAccountService(accountRepo repository.AccountRepository) AccountService {
	return &accountService{
		accountRepo: accountRepo,
	}
}
