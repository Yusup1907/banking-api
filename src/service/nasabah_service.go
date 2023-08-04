package service

import (
	"errors"
	"time"

	"github.com/Yusup1907/banking-api/src/model"
	"github.com/Yusup1907/banking-api/src/repository"
	"github.com/Yusup1907/banking-api/src/utils"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type NasabahService interface {
	RegisterNasabah(nasabah *model.Nasabah) error
	Login(loginRequest *model.LoginRequest, c *gin.Context) (*model.Nasabah, error)
	GetAllNasabah(page int, pageSize int) ([]*model.Nasabah, error)
	GetNasabahById(id string) (*model.Nasabah, error)
	UpdateNasabah(nasabah *model.Nasabah) error
}

type nasabahService struct {
	nasabahRepo repository.NasabahRepository
}

func (s *nasabahService) RegisterNasabah(nasabah *model.Nasabah) error {
	// Check if a nasabah with the same email already exists
	existingNasabah, err := s.nasabahRepo.FindByEmail(nasabah.Email)
	if err != nil {
		return err
	}
	if existingNasabah != nil {
		return errors.New("a nasabah with this email already exists")
	}

	// Generate UUID for the Id field
	nasabah.Id = uuid.New().String()

	if !utils.IsValidPassword(nasabah.Password) {
		return errors.New("Password is not strong enough")
	}

	// Validasi data pengguna
	if !utils.IsValidEmail(nasabah.Email) {
		return errors.New("Invalid email address")
	}
	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(nasabah.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	nasabah.Password = string(hashedPassword)

	// Additional validation or business logic can be performed here before adding the nasabah.

	return s.nasabahRepo.AddNasabah(nasabah)
}

func (s *nasabahService) Login(loginRequest *model.LoginRequest, c *gin.Context) (*model.Nasabah, error) {
	session := sessions.Default(c)

	existSession := session.Get("Email")
	if existSession != nil {
		return nil, errors.New("You are already logged")
	}

	nasabah, err := s.nasabahRepo.FindByEmail(loginRequest.Email)
	if err != nil {
		return nil, err
	}

	if nasabah == nil {
		return nil, errors.New("invalid email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(nasabah.Password), []byte(loginRequest.Password))
	if err != nil {
		return nil, errors.New("invalid password")
	}

	// Set session data
	session.Set("Email", nasabah.Email)
	session.Set("Id", nasabah.Id)
	session.Save()

	// Clear password before returning the nasabah data
	nasabah.Password = ""
	return nasabah, nil
}

func (s *nasabahService) GetAllNasabah(page int, pageSize int) ([]*model.Nasabah, error) {
	if page < 1 {
		return nil, errors.New("invalid page number")
	}
	if pageSize < 1 {
		return nil, errors.New("invalid page size")
	}

	// Panggil fungsi GetAllNasabah dari repository
	nasabahs, err := s.nasabahRepo.GetAllNasabah(page, pageSize)
	if err != nil {
		return nil, err
	}

	return nasabahs, nil
}

func (s *nasabahService) GetNasabahById(id string) (*model.Nasabah, error) {
	return s.nasabahRepo.GetNasabahById(id)
}

func (s *nasabahService) UpdateNasabah(nasabah *model.Nasabah) error {
	existingNasabah, err := s.nasabahRepo.GetNasabahById(nasabah.Id)
	if err != nil {
		return err
	}

	if existingNasabah == nil {
		return errors.New("nasabah not found")
	}

	nasabah.CreatedAt = existingNasabah.CreatedAt
	nasabah.UpdatedAt = time.Now()

	err = s.nasabahRepo.UpdateNasabah(nasabah)
	if err != nil {
		return err
	}

	return nil
}

func NewNasabahService(nasabahRepo repository.NasabahRepository) NasabahService {
	return &nasabahService{
		nasabahRepo: nasabahRepo,
	}
}
