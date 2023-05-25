package models

import "gorm.io/gorm"

type Payment struct {
	gorm.Model
	Type           string `json:"type" form:"type"`
	Name           string `json:"name" form:"name"`
	Account_number string `json:"account_number" form:"account_number"`
}
