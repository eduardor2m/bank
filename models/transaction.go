package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	AccountOriginID      int     `gorm:"type:int"`
	AccountDestinationId int     `gorm:"type:int"`
	Value                float64 `gorm:"type:float" validate:"nonzero"`
}

func (c *Transaction) Validate() error {
	return validator.Validate(c)
}
