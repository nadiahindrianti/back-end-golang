package models

import "gorm.io/gorm"

type HistorySearchStation struct {
	gorm.Model
	UserID               uint `json:"user_id" form:"user_id"`
	StationOriginID      uint `json:"station_origin_id" form:"station_origin_id"`
	StationDestinationID uint `json:"station_destination_id" form:"station_destination_id"`
}
