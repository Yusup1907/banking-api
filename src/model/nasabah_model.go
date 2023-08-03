package model

import "time"

type Nasabah struct {
	Id            string    `json:"id"`
	Name          string    `json:"name"`
	Address       string    `json:"address"`
	PhoneNumber   string    `json:"phone_number"`
	Email         string    `json:"email"`
	Password      string    `json:"password"`
	DateOfBirth   string    `json:"dateOfBirth"`
	Gender        string    `json:"gender"`
	MaritalStatus string    `json:"marital_status"`
	Occupation    string    `json:"occupation"`
	Income        int       `json:"income"`
	Nationality   string    `json:"nationality"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
