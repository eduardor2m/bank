package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	NumberAccountOrigin      string  `gorm:"type:varchar(100)" json:"number_account_origin"`
	NumberAccountDestination string  `gorm:"type:varchar(100)" json:"number_account_destination"`
	Value                    float64 `gorm:"type:float" validate:"nonzero" json:"value"`
	AccountID                uint
}

func (c *Transaction) Validate() error {
	return validator.Validate(c)
}
