package repositories

import (
	"back-end-golang/models"

	"gorm.io/gorm"
)

type HistorySearchStationRepository interface {
	GetAllHistorySearchStation(userID uint, page, limit int) ([]models.HistorySearchStation, int, error)
	GetHistorySearchStationByID(userID, id uint) (models.HistorySearchStation, error)
	CreateHistorySearchStation(HistorySearchStation models.HistorySearchStation) (models.HistorySearchStation, error)
	UpdateHistorySearchStation(HistorySearchStation models.HistorySearchStation) (models.HistorySearchStation, error)
}

type historySearchStationRepository struct {
	db *gorm.DB
}

func NewHistorySearchStationRepository(db *gorm.DB) HistorySearchStationRepository {
	return &historySearchStationRepository{db}
}

func (r *historySearchStationRepository) GetAllHistorySearchStation(userID uint, page, limit int) ([]models.HistorySearchStation, int, error) {
	var (
		historySearchStation []models.HistorySearchStation
		count                 int64
	)
	err := r.db.Where("userID = ?", userID).Find(&historySearchStation).Count(&count).Error
	if err != nil {
		return historySearchStation, 0, err
	}

	offset := (page - 1) * limit

	err = r.db.Limit(limit).Offset(offset).Find(&historySearchStation).Error

	return historySearchStation, int(count), err
}

func (r *historySearchStationRepository) GetHistorySearchStationByID(userID ,id uint) (models.HistorySearchStation, error) {
	var historySearchStation models.HistorySearchStation
	err := r.db.Where("id = ? AND user_ID = ?", id, userID).First(&historySearchStation).Error
	return historySearchStation, err
}

func (r *historySearchStationRepository) CreateHistorySearchStation(historySearchStation models.HistorySearchStation) (models.HistorySearchStation, error) {
	err := r.db.Create(&historySearchStation).Error
	return historySearchStation, err
}

func (r *historySearchStationRepository) UpdateHistorySearchStation(historySearchStation models.HistorySearchStation) (models.HistorySearchStation, error) {
	err := r.db.Save(&historySearchStation).Error
	return historySearchStation, err
}