// file: service/transaction_service.go
package service

import (
	"errors"

	"github.com/Yusup1907/banking-api/src/model"
	"github.com/Yusup1907/banking-api/src/repository"
)

type TransactionService interface {
	Transfer(transaction *model.Transaction) error
}

type transactionService struct {
	accountRepo     repository.AccountRepository
	transactionRepo repository.TransactionRepository
}

func (s *transactionService) Transfer(transaction *model.Transaction) error {
	// Validate transaction data
	if err := transaction.Validate(); err != nil {
		return err
	}

	// Get sender account from the account repository
	senderAccount, err := s.accountRepo.GetAccountByID(transaction.AccountID)
	if err != nil {
		return errors.New("failed to get sender account")
	}

	// Get receiver account from the account repository
	receiverAccount, err := s.accountRepo.GetAccountByAccountNumber(senderAccount.AccountNumber)
	if err != nil {
		return errors.New("failed to get receiver account")
	}

	if receiverAccount == nil {
		return errors.New("receiver account not found")
	}

	// Check if the sender account has enough balance
	if senderAccount.Balance < transaction.Amount {
		return errors.New("insufficient balance")
	}

	// Perform transfer
	senderAccount.Balance -= transaction.Amount
	receiverAccount.Balance += transaction.Amount

	// Update account data in the account repository
	if err := s.accountRepo.UpdateAccount(senderAccount); err != nil {
		return errors.New("failed to update sender account")
	}

	if err := s.accountRepo.UpdateAccount(receiverAccount); err != nil {
		return errors.New("failed to update receiver account")
	}

	// Save the transaction data in the transaction repository
	if err := s.transactionRepo.AddTransaction(transaction); err != nil {
		return errors.New("failed to save transaction")
	}

	return nil
}

func NewTransactionService(accountRepo repository.AccountRepository, transactionRepo repository.TransactionRepository) TransactionService {
	return &transactionService{
		accountRepo:     accountRepo,
		transactionRepo: transactionRepo,
	}
}
