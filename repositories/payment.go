package repositories

import (
	"back-end-golang/models"

	"gorm.io/gorm"
)

type PaymentRepository interface {
	GetPaymentByID(id uint) (models.Payment, error)
	CreatePayment(payment models.Payment) (models.Payment, error)
	UpdatePayment(payment models.Payment) (models.Payment, error)
	DeletePayment(payment models.Payment) error
}

type paymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) PaymentRepository {
	return &paymentRepository{db}
}

// Implementasi fungsi-fungsi dari interface ItemRepository

func (r *paymentRepository) GetPaymentByID(id uint) (models.Payment, error) {
	var payment models.Payment
	err := r.db.Where("id = ?", id).First(&payment).Error
	return payment, err
}

func (r *paymentRepository) CreatePayment(payment models.Payment) (models.Payment, error) {
	err := r.db.Create(&payment).Error
	return payment, err
}

func (r *paymentRepository) UpdatePayment(payment models.Payment) (models.Payment, error) {
	err := r.db.Save(&payment).Error
	return payment, err
}

func (r *paymentRepository) DeletePayment(payment models.Payment) error {
	err := r.db.Delete(&payment).Error
	return err
}
