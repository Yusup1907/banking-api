package repository

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"time"

	"github.com/Yusup1907/banking-api/src/model"
)

type NasabahRepository interface {
	AddNasabah(nasabah *model.Nasabah) error
	FindByEmail(email string) (*model.Nasabah, error)
}

type nasabahRepository struct {
	filePath string
}

func (n *nasabahRepository) AddNasabah(nasabah *model.Nasabah) error {
	nasabahs, err := readNasabahFromFile(n.filePath)
	if err != nil {
		return err
	}

	nasabah.CreatedAt = time.Now()
	nasabah.UpdatedAt = time.Now()
	nasabahs = append(nasabahs, nasabah)

	err = saveNasabahToFile(n.filePath, nasabahs)
	if err != nil {
		return err
	}

	return nil
}

func (n *nasabahRepository) FindByEmail(email string) (*model.Nasabah, error) {
	nasabahs, err := readNasabahFromFile(n.filePath)
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

func readNasabahFromFile(filePath string) ([]*model.Nasabah, error) {
	file, err := os.OpenFile(filePath, os.O_RDONLY|os.O_CREATE, 0644)
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

func saveNasabahToFile(filePath string, nasabahs []*model.Nasabah) error {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
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

func NewNasabahRepository(filePath string) NasabahRepository {
	return &nasabahRepository{
		filePath: filePath,
	}

}
