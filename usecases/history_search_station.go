package usecases

import (
	"back-end-golang/dtos"
	"back-end-golang/models"
	"back-end-golang/repositories"
)

type HistorySearchStationUseCase interface {
	GetAllHistorySearchStation(userID uint, page, limit int) ([]dtos.HistorySearchStationResponse, int, error)
	GetHistorySearchStationByID(userID, id uint) (dtos.HistorySearchStationResponse, error)
	CreateHistorySearchStation(userID uint, historySearchStationInput dtos.HistorySearchStationInput) (dtos.HistorySearchStationResponse, error)
	UpdateHistorySearchStation(userID, id uint, historySearchStationInput dtos.HistorySearchStationInput) (dtos.HistorySearchStationResponse, error)
}

type historySearchStationUsecase struct {
	historySearchStationRepo repositories.HistorySearchStationRepository
	userRepo                 repositories.UserRepository
}

func NewHistorySearchStationUsecase(historySearchStationRepo repositories.HistorySearchStationRepository, userRepo repositories.UserRepository) HistorySearchStationUseCase {
	return &historySearchStationUsecase{historySearchStationRepo, userRepo}
}

//

func (u *historySearchStationUsecase) GetAllHistorySearchStation(userID uint, page, limit int) ([]dtos.HistorySearchStationResponse, int, error) {
	var historySearchStationResponses []dtos.HistorySearchStationResponse

	historySearchStations, count, err := u.historySearchStationRepo.GetAllHistorySearchStation(userID, page, limit)
	if err != nil {
		return historySearchStationResponses, count, err
	}

	for _, historySearchStation := range historySearchStations {
		historySearchStationResponse := dtos.HistorySearchStationResponse{
			ID:                   historySearchStation.ID,
			UserID:               historySearchStation.UserID,
			StationOriginID:      historySearchStation.StationOriginID,
			StationDestinationID: historySearchStation.StationDestinationID,
			CreatedAt:            &historySearchStation.CreatedAt,
			UpdatedAt:            &historySearchStation.UpdatedAt,
		}
		historySearchStationResponses = append(historySearchStationResponses, historySearchStationResponse)
	}

	return historySearchStationResponses, count, nil

}

func (u *historySearchStationUsecase) GetHistorySearchStationByID(userID, id uint) (dtos.HistorySearchStationResponse, error) {
	var historySearchStationResponses dtos.HistorySearchStationResponse

	historySearchStation, err := u.historySearchStationRepo.GetHistorySearchStationByID(userID, id)
	if err != nil {
		return historySearchStationResponses, err
	}

	historySearchStationResponse := dtos.HistorySearchStationResponse{
		ID:                   historySearchStation.ID,
		UserID:               historySearchStation.UserID,
		StationOriginID:      historySearchStation.StationOriginID,
		StationDestinationID: historySearchStation.StationDestinationID,
		CreatedAt:            &historySearchStation.CreatedAt,
		UpdatedAt:            &historySearchStation.UpdatedAt,
	}

	return historySearchStationResponse, nil
}

func (u *historySearchStationUsecase) CreateHistorySearchStation(userID uint, historySearchStationInput dtos.HistorySearchStationInput) (dtos.HistorySearchStationResponse, error) {
	var (
		historySearchStation         models.HistorySearchStation
		historySearchStationResponse dtos.HistorySearchStationResponse
	)

	historySearchStation.UserID = userID
	historySearchStation.StationOriginID = historySearchStationInput.StationOriginID
	historySearchStation.StationDestinationID = historySearchStationInput.StationDestinationID

	historySearchStation, err := u.historySearchStationRepo.CreateHistorySearchStation(historySearchStation)
	if err != nil {
		return historySearchStationResponse, err
	}

	historySearchStationResponse.ID = historySearchStation.ID
	historySearchStationResponse.UserID = historySearchStation.UserID
	historySearchStationResponse.StationOriginID = historySearchStation.StationOriginID
	historySearchStationResponse.StationDestinationID = historySearchStation.StationDestinationID
	historySearchStationResponse.CreatedAt = &historySearchStation.CreatedAt
	historySearchStationResponse.UpdatedAt = &historySearchStation.UpdatedAt

	return historySearchStationResponse, nil

}

func (u *historySearchStationUsecase) UpdateHistorySearchStation(userID, id uint, historySearchStationInput dtos.HistorySearchStationInput) (dtos.HistorySearchStationResponse, error) {
	var historySearchStation models.HistorySearchStation
	var historySearchStationResponse dtos.HistorySearchStationResponse

	historySearchStation, err := u.historySearchStationRepo.GetHistorySearchStationByID(userID, id)
	if err != nil {
		return historySearchStationResponse, err
	}

	historySearchStation.UserID = userID
	historySearchStation.StationOriginID = historySearchStationInput.StationOriginID
	historySearchStation.StationDestinationID = historySearchStationInput.StationDestinationID

	historySearchStation, err = u.historySearchStationRepo.UpdateHistorySearchStation(historySearchStation)

	if err != nil {
		return historySearchStationResponse, err
	}

	historySearchStationResponse.ID = userID
	historySearchStationResponse.UserID = historySearchStation.UserID
	historySearchStationResponse.StationOriginID = historySearchStation.StationOriginID
	historySearchStationResponse.StationDestinationID = historySearchStation.StationDestinationID
	historySearchStationResponse.CreatedAt = &historySearchStation.CreatedAt
	historySearchStationResponse.UpdatedAt = &historySearchStation.UpdatedAt

	return historySearchStationResponse, nil
}
