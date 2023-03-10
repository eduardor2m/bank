package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	AccountNumber string  `gorm:"type:varchar(100);unique_index" json:"account_number"`
	AgencyNumber  string  `gorm:"type:varchar(100)" json:"agency_number"`
	Balance       float64 `gorm:"type:float" json:"balance"`
	Transactions  []Transaction
	ClientID      uint
}

func (a *Account) Validate() error {
	return validator.Validate(a)
}
