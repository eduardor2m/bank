package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Client struct {
	gorm.Model
	Name     string `gorm:"type:varchar(100)" validate:"nonzero" json:"name"`
	CPF      string `gorm:"type:varchar(100);unique_index" validate:"nonzero" json:"cpf"`
	Email    string `gorm:"type:varchar(100);unique_index" validate:"nonzero" json:"email"`
	Password string `gorm:"type:varchar(100)" validate:"nonzero" json:"password"`
	Account  Account
}

func (c *Client) Validate() error {
	return validator.Validate(c)
}
