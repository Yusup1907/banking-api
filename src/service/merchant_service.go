package service

import (
	"github.com/Yusup1907/banking-api/src/model"
	"github.com/Yusup1907/banking-api/src/repository"
)

type MerchantService interface {
	AddMerchant(merchant *model.Merchant) error
	GetAllMerchants() ([]*model.Merchant, error)
	GetMerchantByID(id string) (*model.Merchant, error)
	UpdateMerchant(merchant *model.Merchant) error
	DeleteMerchant(id string) error
}

type merchantService struct {
	merchantRepo repository.MerchantRepository
}

func (s *merchantService) AddMerchant(merchant *model.Merchant) error {
	return s.merchantRepo.AddMerchant(merchant)
}

func (s *merchantService) GetAllMerchants() ([]*model.Merchant, error) {
	return s.merchantRepo.GetAllMerchants()
}

func (s *merchantService) GetMerchantByID(id string) (*model.Merchant, error) {
	return s.merchantRepo.GetMerchantByID(id)
}

func (s *merchantService) UpdateMerchant(merchant *model.Merchant) error {
	return s.merchantRepo.UpdateMerchant(merchant)
}

func (s *merchantService) DeleteMerchant(id string) error {
	return s.merchantRepo.DeleteMerchant(id)
}

func NewMerchantService(merchantRepo repository.MerchantRepository) MerchantService {
	return &merchantService{
		merchantRepo: merchantRepo,
	}
}
