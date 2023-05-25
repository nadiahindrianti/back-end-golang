package dtos

import "time"

type PaymentInput struct {
	Type           string `json:"type" form:"type" `
	Name           string `json:"name" form:"name"`
	Account_number string `json:"account_number" form:"account_number" `
}

type PaymentResponse struct {
	PaymentID      uint      `json:"payment_id"`
	Type           string    `json:"type" form:"type" `
	Name           string    `json:"name" form:"name"`
	Account_number string    `json:"account_number" form:"account_number" `
	CreatedAt      time.Time `json:"created_at" example:"2023-05-17T15:07:16.504+07:00"`
	UpdatedAt      time.Time `json:"updated_at" example:"2023-05-17T15:07:16.504+07:00"`
}
