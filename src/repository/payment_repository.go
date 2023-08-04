package repository

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/Yusup1907/banking-api/src/model"
)

const (
	transactionFilePath = "../data/payment.json"
)

type TransactionRepository interface {
	AddTransaction(transaction *model.Transaction) error
	GetAllTransactions() ([]*model.Transaction, error)
}

type transactionRepository struct{}

func (t *transactionRepository) AddTransaction(transaction *model.Transaction) error {
	transactions, err := t.readTransactionsFromFile()
	if err != nil {
		return err
	}

	transactions = append(transactions, transaction)

	err = t.saveTransactionsToFile(transactions)
	if err != nil {
		return err
	}

	return nil
}

func (t *transactionRepository) GetAllTransactions() ([]*model.Transaction, error) {
	return t.readTransactionsFromFile()
}

func (t *transactionRepository) readTransactionsFromFile() ([]*model.Transaction, error) {
	file, err := os.OpenFile(transactionFilePath, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var transactions []*model.Transaction
	if len(data) > 0 {
		err = json.Unmarshal(data, &transactions)
		if err != nil {
			return nil, err
		}
	}

	return transactions, nil
}

func (t *transactionRepository) saveTransactionsToFile(transactions []*model.Transaction) error {
	file, err := os.OpenFile(transactionFilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := json.MarshalIndent(transactions, "", "  ")
	if err != nil {
		return err
	}

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func NewTransactionRepository() TransactionRepository {
	return &transactionRepository{}
}
