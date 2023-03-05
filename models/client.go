package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Client struct {
	gorm.Model
	Name     string `gorm:"type:varchar(100)" validate:"nonzero"`
	CPF      string `gorm:"type:varchar(100);unique_index" validate:"nonzero"`
	Email    string `gorm:"type:varchar(100);unique_index" validate:"nonzero"`
	Password string `gorm:"type:varchar(100)" validate:"nonzero"`
}

func (c *Client) Validate() error {
	return validator.Validate(c)
}
