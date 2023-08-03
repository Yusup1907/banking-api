package model

type Account struct {
	Id            string `json:"id"`
	NasabahId     string `json:"NasabahId"`
	AccountNumber string `json:"account_number"`
	Balance       int    `json:"balance"`
	Status        string `json:"status"`
}
