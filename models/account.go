package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	AccountNumber string  `gorm:"type:varchar(100);unique_index" validate:"nonzero"`
	AgencyNumber  string  `gorm:"type:varchar(100)" validate:"nonzero"`
	Balance       float64 `gorm:"type:float" validate:"nonzero"`
}

func (a *Account) Validate() error {
	return validator.Validate(a)
}
