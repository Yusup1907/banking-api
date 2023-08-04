package repository

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"

	"github.com/Yusup1907/banking-api/src/model"
	"github.com/google/uuid"
)

const (
	accountFilePath = "../data/account.json"
)

type AccountRepository interface {
	AddAccount(account *model.Account) error
	GetAccountByID(accountID string) (*model.Account, error)
	GetAccountByAccountNumber(accountNumber string) (*model.Account, error)
	UpdateAccount(account *model.Account) error
}

type accountRepository struct{}

func (a *accountRepository) AddAccount(account *model.Account) error {
	accounts, err := a.readAccountsFromFile()
	if err != nil {
		return err
	}

	// Generate a unique ID for the new account (you can use a UUID library or any other method)
	account.Id = uuid.New().String()

	accounts = append(accounts, account)

	err = a.saveAccountsToFile(accounts)
	if err != nil {
		return err
	}

	return nil
}

func (a *accountRepository) GetAccountByID(accountID string) (*model.Account, error) {
	accounts, err := a.readAccountsFromFile()
	if err != nil {
		return nil, err
	}

	for _, account := range accounts {
		if account.Id == accountID {
			return account, nil
		}
	}

	return nil, nil // Account not found
}

func (a *accountRepository) GetAccountByAccountNumber(accountNumber string) (*model.Account, error) {
	accounts, err := a.readAccountsFromFile()
	if err != nil {
		return nil, err
	}

	for _, account := range accounts {
		if account.AccountNumber == accountNumber {
			return account, nil
		}
	}

	return nil, nil // Account not found
}

func (a *accountRepository) UpdateAccount(account *model.Account) error {
	accounts, err := a.readAccountsFromFile()
	if err != nil {
		return err
	}

	found := false
	for i, existingAccount := range accounts {
		if existingAccount.Id == account.Id {
			accounts[i] = account
			found = true
			break
		}
	}

	if !found {
		return errors.New("account not found")
	}

	err = a.saveAccountsToFile(accounts)
	if err != nil {
		return err
	}

	return nil
}

func (a *accountRepository) readAccountsFromFile() ([]*model.Account, error) {
	file, err := os.OpenFile(accountFilePath, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var accounts []*model.Account
	if len(data) > 0 {
		err = json.Unmarshal(data, &accounts)
		if err != nil {
			return nil, err
		}
	}

	return accounts, nil
}

func (a *accountRepository) saveAccountsToFile(accounts []*model.Account) error {
	file, err := os.OpenFile(accountFilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := json.MarshalIndent(accounts, "", "  ")
	if err != nil {
		return err
	}

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func NewAccountRepository() AccountRepository {
	return &accountRepository{}
}
