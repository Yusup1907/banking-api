package service

import (
	"errors"
	"time"

	"github.com/Yusup1907/banking-api/src/model"
	"github.com/Yusup1907/banking-api/src/repository"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type NasabahService interface {
	RegisterNasabah(nasabah *model.Nasabah) error
	Login(email, password string) (string, error)
}

type nasabahService struct {
	nasabahRepo repository.NasabahRepository
	secretKey   string
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

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(nasabah.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	nasabah.Password = string(hashedPassword)

	// Additional validation or business logic can be performed here before adding the nasabah.

	return s.nasabahRepo.AddNasabah(nasabah)
}

func (s *nasabahService) Login(email, password string) (string, error) {
	nasabah, err := s.nasabahRepo.FindByEmail(email)
	if err != nil {
		return "", err
	}

	if nasabah == nil {
		return "", errors.New("invalid email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(nasabah.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid password")
	}

	// Generate a new JWT token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = nasabah.Email
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func NewNasabahService(nasabahRepo repository.NasabahRepository, secretKey string) NasabahService {
	return &nasabahService{
		nasabahRepo: nasabahRepo,
		secretKey:   secretKey,
	}
}
