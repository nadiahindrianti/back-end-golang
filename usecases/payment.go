package usecases

import (
	"back-end-golang/dtos"
	"back-end-golang/models"
	"back-end-golang/repositories"
)

type PaymentUsecase interface {
	GetPaymentByID(id uint) (dtos.PaymentResponse, error)
	CreatePayment(paymentInput *dtos.PaymentInput) (dtos.PaymentResponse, error)
	UpdatePayment(id uint, paymentInput dtos.PaymentInput) (dtos.PaymentResponse, error)
	DeletePayment(id uint) error
}

type paymentUsecase struct {
	paymentRepo repositories.PaymentRepository
}

func NewPaymentUsecase(PaymentRepo repositories.PaymentRepository) PaymentUsecase {
	return &paymentUsecase{PaymentRepo}
}

func (u *paymentUsecase) GetPaymentByID(id uint) (dtos.PaymentResponse, error) {
	var paymentResponses dtos.PaymentResponse
	payment, err := u.paymentRepo.GetPaymentByID(id)
	if err != nil {
		return paymentResponses, err
	}
	paymentResponse := dtos.PaymentResponse{
		PaymentID:      payment.ID,
		Type:           payment.Type,
		Name:           payment.Name,
		Account_number: payment.Account_number,
		CreatedAt:      payment.CreatedAt,
		UpdatedAt:      payment.UpdatedAt,
	}
	return paymentResponse, nil
}

func (u *paymentUsecase) CreatePayment(paymentInput *dtos.PaymentInput) (dtos.PaymentResponse, error) {
	var paymentResponses dtos.PaymentResponse
	CreatePayment := models.Payment{
		Type:           paymentInput.Type,
		Name:           paymentInput.Name,
		Account_number: paymentInput.Account_number,
	}

	createdPayment, err := u.paymentRepo.CreatePayment(CreatePayment)
	if err != nil {
		return paymentResponses, err
	}

	paymentResponse := dtos.PaymentResponse{
		PaymentID:      createdPayment.ID,
		Type:           createdPayment.Type,
		Name:           createdPayment.Name,
		Account_number: createdPayment.Account_number,
		CreatedAt:      createdPayment.CreatedAt,
		UpdatedAt:      createdPayment.UpdatedAt,
	}
	return paymentResponse, nil
}

func (u *paymentUsecase) UpdatePayment(id uint, paymentInput dtos.PaymentInput) (dtos.PaymentResponse, error) {
	var payment models.Payment
	var paymentResponse dtos.PaymentResponse

	payment, err := u.paymentRepo.GetPaymentByID(id)
	if err != nil {
		return paymentResponse, err
	}

	payment.Type = paymentInput.Type
	payment.Name = paymentInput.Name
	payment.Account_number = paymentInput.Account_number

	payment, err = u.paymentRepo.UpdatePayment(payment)

	if err != nil {
		return paymentResponse, err
	}

	paymentResponse.PaymentID = payment.ID
	paymentResponse.Type = payment.Type
	paymentResponse.Name = payment.Name
	paymentResponse.Account_number = payment.Account_number
	paymentResponse.CreatedAt = payment.CreatedAt
	paymentResponse.UpdatedAt = payment.UpdatedAt

	return paymentResponse, nil

}

func (u *paymentUsecase) DeletePayment(id uint) error {
	payment, err := u.paymentRepo.GetPaymentByID(id)

	if err != nil {
		return nil
	}

	err = u.paymentRepo.DeletePayment(payment)
	return err
}
