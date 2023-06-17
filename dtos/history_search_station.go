package dtos

import "time"

type HistorySearchStationInput struct {
	StationOriginID      uint `json:"station_origin_id" form:"station_origin_id"`
	StationDestinationID uint `json:"station_destination_id" form:"station_destination_id"`
}

type HistorySearchStationResponse struct {
	ID                   uint       `json:"history_search_station_id" form:"history_search_station_id"`
	UserID               uint       `json:"user_id" form:"user_id"`
	StationOriginID      uint       `json:"station_origin_id" form:"station_origin_id"`
	StationDestinationID uint       `json:"station_destination_id" form:"station_destination_id"`
	CreatedAt            *time.Time `json:"created_at" example:"2023-05-17T15:07:16.504+07:00"`
	UpdatedAt            *time.Time `json:"updated_at" example:"2023-05-17T15:07:16.504+07:00"`
}
