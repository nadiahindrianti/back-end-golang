package repositories

import (
	"back-end-golang/models"
	"time"

	"gorm.io/gorm"
)

type DashboardRepository interface {
	DashboardGetAll() (int, int, int, int, int, int, int, int, []models.TicketTravelerDetail, []models.User, []models.HotelOrder, int, int, error)
}

type dashboardRepository struct {
	db *gorm.DB
}

func NewDashboardRepository(db *gorm.DB) DashboardRepository {
	return &dashboardRepository{db}
}

func (r *dashboardRepository) DashboardGetAll() (int, int, int, int, int, int, int, int, []models.TicketTravelerDetail, []models.User, []models.HotelOrder, int, int, error) {
	var countUser int64
	var countUserToday int64

	user := models.User{}
	err := r.db.Unscoped().Where("role = 'user'").Order("id DESC").Find(&user).Count(&countUser).Error
	if err != nil {
		return 0, 0, 0, 0, 0, 0, 0, 0, []models.TicketTravelerDetail{}, []models.User{}, []models.HotelOrder{}, 0, 0, err
	}

	newUser := []models.User{}
	err = r.db.Unscoped().Where("role = 'user'").Order("id DESC").Limit(10).Find(&newUser).Error
	if err != nil {
		return 0, 0, 0, 0, 0, 0, 0, 0, []models.TicketTravelerDetail{}, []models.User{}, []models.HotelOrder{}, 0, 0, err
	}

	// Get the start and end of the current day
	startOfDay := time.Now().Truncate(24 * time.Hour)
	endOfDay := startOfDay.Add(24 * time.Hour)

	// Count the number of users created today
	err = r.db.Model(&models.User{}).Unscoped().Where("role = 'user' AND created_at BETWEEN ? AND ?", startOfDay, endOfDay).Count(&countUserToday).Error
	if err != nil {
		return 0, 0, 0, 0, 0, 0, 0, 0, []models.TicketTravelerDetail{}, []models.User{}, []models.HotelOrder{}, 0, 0, err
	}

	var countTrain int64
	var countTrainToday int64
	train := models.Train{}
	err = r.db.Unscoped().Find(&train).Count(&countTrain).Error
	if err != nil {
		return 0, 0, 0, 0, 0, 0, 0, 0, []models.TicketTravelerDetail{}, []models.User{}, []models.HotelOrder{}, 0, 0, err
	}

	// Count the number of users created today
	err = r.db.Model(&models.Train{}).Unscoped().Where("created_at BETWEEN ? AND ?", startOfDay, endOfDay).Count(&countTrainToday).Error
	if err != nil {
		return 0, 0, 0, 0, 0, 0, 0, 0, []models.TicketTravelerDetail{}, []models.User{}, []models.HotelOrder{}, 0, 0, err
	}

	var countHotel int64
	var countHotelToday int64
	hotel := models.Hotel{}
	err = r.db.Unscoped().Find(&hotel).Count(&countHotel).Error
	if err != nil {
		return 0, 0, 0, 0, 0, 0, 0, 0, []models.TicketTravelerDetail{}, []models.User{}, []models.HotelOrder{}, 0, 0, err
	}

	// Count the number of users created today
	err = r.db.Model(&models.Hotel{}).Unscoped().Where("created_at BETWEEN ? AND ?", startOfDay, endOfDay).Count(&countHotelToday).Error
	if err != nil {
		return 0, 0, 0, 0, 0, 0, 0, 0, []models.TicketTravelerDetail{}, []models.User{}, []models.HotelOrder{}, 0, 0, err
	}

	ticketOrder := []models.TicketTravelerDetail{}

	var countTicketOrder int64
	var countTicketOrderToday int64
	err = r.db.Find(&ticketOrder).Count(&countTicketOrder).Error
	if err != nil {
		return 0, 0, 0, 0, 0, 0, 0, 0, []models.TicketTravelerDetail{}, []models.User{}, []models.HotelOrder{}, 0, 0, err
	}
	newTicketOrder := []models.TicketTravelerDetail{}
	err = r.db.Order("id DESC").Limit(10).Find(&newTicketOrder).Error

	if err != nil {
		return 0, 0, 0, 0, 0, 0, 0, 0, []models.TicketTravelerDetail{}, []models.User{}, []models.HotelOrder{}, 0, 0, err
	}

	// Count the number of users created today
	err = r.db.Model(&models.TicketOrder{}).Where("created_at BETWEEN ? AND ?", startOfDay, endOfDay).Count(&countTicketOrderToday).Error
	if err != nil {
		return 0, 0, 0, 0, 0, 0, 0, 0, []models.TicketTravelerDetail{}, []models.User{}, []models.HotelOrder{}, 0, 0, err
	}

	hotelOrder := []models.HotelOrder{}

	var countHotelOrder int64
	var countHotelOrderToday int64
	err = r.db.Find(&hotelOrder).Count(&countHotelOrder).Error
	if err != nil {
		return 0, 0, 0, 0, 0, 0, 0, 0, []models.TicketTravelerDetail{}, []models.User{}, []models.HotelOrder{}, 0, 0, err
	}
	newHotelOrder := []models.HotelOrder{}
	err = r.db.Order("id DESC").Limit(10).Find(&newHotelOrder).Error

	if err != nil {
		return 0, 0, 0, 0, 0, 0, 0, 0, []models.TicketTravelerDetail{}, []models.User{}, []models.HotelOrder{}, 0, 0, err
	}

	// Count the number of users created today
	err = r.db.Model(&models.HotelOrder{}).Where("created_at BETWEEN ? AND ?", startOfDay, endOfDay).Count(&countHotelOrderToday).Error
	if err != nil {
		return 0, 0, 0, 0, 0, 0, 0, 0, []models.TicketTravelerDetail{}, []models.User{}, []models.HotelOrder{}, 0, 0, err
	}

	return int(countUser), int(countUserToday), int(countTrain), int(countTrainToday), int(countHotel), int(countHotelToday),  int(countTicketOrder), int(countTicketOrderToday), newTicketOrder, newUser, newHotelOrder, int(countHotelOrder), int(countHotelOrderToday), nil
}
