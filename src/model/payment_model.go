// file: model/transaction.go
package model

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	Id              string    `json:"id"`
	AccountID       string    `json:"accountId"`
	MerchantID      string    `json:"merchantId"`
	Amount          int       `json:"amount"`
	TransactionType string    `json:"transaction_type"`
	TransactionTime time.Time `json:"transaction_datetime"`
}

func NewTransaction(accountID, merchantID string, amount int, transactionType string, transactionTime time.Time) *Transaction {
	return &Transaction{
		Id:              uuid.New().String(),
		AccountID:       accountID,
		MerchantID:      merchantID,
		Amount:          amount,
		TransactionType: transactionType,
		TransactionTime: transactionTime,
	}
}

func (t *Transaction) Validate() error {
	// Check if AccountID is a valid UUID
	_, err := uuid.Parse(t.AccountID)
	if err != nil {
		return errors.New("invalid AccountID")
	}

	// Check if MerchantID is a valid UUID
	_, err = uuid.Parse(t.MerchantID)
	if err != nil {
		return errors.New("invalid MerchantID")
	}

	if t.Amount <= 0 {
		return errors.New("amount must be a positive value")
	}

	// Add other validations as needed...

	return nil
}
