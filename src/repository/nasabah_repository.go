package repository

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"time"

	"github.com/Yusup1907/banking-api/src/model"
)

const nasabahFilePath = "../data/nasabah.json" // File path tetap ada dalam package repository

type NasabahRepository interface {
	AddNasabah(nasabah *model.Nasabah) error
	FindByEmail(email string) (*model.Nasabah, error)
	GetAllNasabah(page int, pageSize int) ([]*model.Nasabah, error)
	GetNasabahById(id string) (*model.Nasabah, error)
}

type nasabahRepository struct{}

func (n *nasabahRepository) AddNasabah(nasabah *model.Nasabah) error {
	nasabahs, err := readNasabahFromFile()
	if err != nil {
		return err
	}

	nasabah.CreatedAt = time.Now()
	nasabah.UpdatedAt = time.Now()
	nasabahs = append(nasabahs, nasabah)

	err = saveNasabahToFile(nasabahs)
	if err != nil {
		return err
	}

	return nil
}

func (n *nasabahRepository) FindByEmail(email string) (*model.Nasabah, error) {
	nasabahs, err := readNasabahFromFile()
	if err != nil {
		return nil, err
	}

	for _, nasabah := range nasabahs {
		if nasabah.Email == email {
			return nasabah, nil
		}
	}

	return nil, nil // Return nil if the nasabah with the given email is not found
}

func (n *nasabahRepository) GetAllNasabah(page int, pageSize int) ([]*model.Nasabah, error) {
	nasabahs, err := readNasabahFromFile()
	if err != nil {
		return nil, err
	}

	startIndex := (page - 1) * pageSize
	endIndex := startIndex + pageSize

	if startIndex >= len(nasabahs) {
		return []*model.Nasabah{}, nil
	}

	if endIndex > len(nasabahs) {
		endIndex = len(nasabahs)
	}

	pagedNasabahs := nasabahs[startIndex:endIndex]

	return pagedNasabahs, nil
}

func (n *nasabahRepository) GetNasabahById(id string) (*model.Nasabah, error) {
	nasabahs, err := readNasabahFromFile()
	if err != nil {
		return nil, err
	}

	for _, nasabah := range nasabahs {
		if nasabah.Id == id {
			return nasabah, nil
		}
	}

	return nil, nil // Return nil if the nasabah with the given ID is not found
}

func readNasabahFromFile() ([]*model.Nasabah, error) {
	file, err := os.OpenFile(nasabahFilePath, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var nasabahs []*model.Nasabah
	if len(data) > 0 {
		err = json.Unmarshal(data, &nasabahs)
		if err != nil {
			return nil, err
		}
	}

	return nasabahs, nil
}

func saveNasabahToFile(nasabahs []*model.Nasabah) error {
	file, err := os.OpenFile(nasabahFilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := json.MarshalIndent(nasabahs, "", "  ")
	if err != nil {
		return err
	}

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func NewNasabahRepository() NasabahRepository {
	return &nasabahRepository{}
}
