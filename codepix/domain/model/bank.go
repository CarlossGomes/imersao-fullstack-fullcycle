package model

import (
	"time"
	uuid "github.com/satori/go.uuid"
)

func (bank *Bank) isValid() error {
	return nil
}

type Bank struct {
	ID	 string `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewBank(code string, name string) (*Bank, error){
	bank := Bank{
		Code: code,
		Name: name,
	}

	bank.ID = uuid.NewV4().String()
	bank.CreatedAt = time.Now()

	err := bank.isValid()
	if err != nil {
		return nil, err
	}

	return &bank, nil

}