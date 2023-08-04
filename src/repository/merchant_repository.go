package repository

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"

	"github.com/Yusup1907/banking-api/src/model"
	"github.com/google/uuid"
)

const merchantFilePath = "../data/merchant.json"

type MerchantRepository interface {
	AddMerchant(merchant *model.Merchant) error
	GetAllMerchants() ([]*model.Merchant, error)
	GetMerchantByID(id string) (*model.Merchant, error)
	UpdateMerchant(merchant *model.Merchant) error
	DeleteMerchant(id string) error
}

type merchantRepository struct{}

func (m *merchantRepository) AddMerchant(merchant *model.Merchant) error {
	merchants, err := readMerchantsFromFile()
	if err != nil {
		return err
	}

	merchant.Id = uuid.New().String()
	// Set the merchant creation timestamp
	merchants = append(merchants, merchant)

	err = saveMerchantsToFile(merchants)
	if err != nil {
		return err
	}

	return nil
}

func (m *merchantRepository) GetAllMerchants() ([]*model.Merchant, error) {
	return readMerchantsFromFile()
}

func (m *merchantRepository) GetMerchantByID(id string) (*model.Merchant, error) {
	merchants, err := readMerchantsFromFile()
	if err != nil {
		return nil, err
	}

	for _, merchant := range merchants {
		if merchant.Id == id {
			return merchant, nil
		}
	}

	return nil, nil
}

func (m *merchantRepository) UpdateMerchant(merchant *model.Merchant) error {
	merchants, err := readMerchantsFromFile()
	if err != nil {
		return err
	}

	found := false
	for i, existingMerchant := range merchants {
		if existingMerchant.Id == merchant.Id {
			merchants[i] = merchant
			found = true
			break
		}
	}

	if !found {
		return errors.New("merchant not found")
	}

	err = saveMerchantsToFile(merchants)
	if err != nil {
		return err
	}

	return nil
}

func (m *merchantRepository) DeleteMerchant(id string) error {
	merchants, err := readMerchantsFromFile()
	if err != nil {
		return err
	}

	found := false
	for i, merchant := range merchants {
		if merchant.Id == id {
			// Remove the merchant from the slice
			merchants = append(merchants[:i], merchants[i+1:]...)
			found = true
			break
		}
	}

	if !found {
		return errors.New("merchant not found")
	}

	err = saveMerchantsToFile(merchants)
	if err != nil {
		return err
	}

	return nil
}

func readMerchantsFromFile() ([]*model.Merchant, error) {
	file, err := os.OpenFile(merchantFilePath, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var merchants []*model.Merchant
	if len(data) > 0 {
		err = json.Unmarshal(data, &merchants)
		if err != nil {
			return nil, err
		}
	}

	return merchants, nil
}

func saveMerchantsToFile(merchants []*model.Merchant) error {
	file, err := os.OpenFile(merchantFilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := json.MarshalIndent(merchants, "", "  ")
	if err != nil {
		return err
	}

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func NewMerchantRepository() MerchantRepository {
	return &merchantRepository{}
}
