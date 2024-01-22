package dtos

import "time"

type TrainCarriageInput struct {
	TrainID uint   `json:"train_id" form:"train_id" example:"1"`
	Class   string `json:"class" form:"class" example:"Ekonomi"`
	Name    string `json:"name" form:"name" example:"Gerbong 1"`
	Price   int    `json:"price" form:"Price" example:"50000"`
}

type TrainCarriageResponse struct {
	TrainCarriageID uint                `json:"train_carriage_id" example:"1"`
	Train           TrainResponse       `json:"train,omitempty"`
	Name            string              `json:"name" example:"Gerbong 1"`
	Seat            []TrainSeatResponse `json:"seat,omitempty"`
	CreatedAt       time.Time           `json:"created_at" example:"2023-05-17T15:07:16.504+07:00"`
	UpdatedAt       time.Time           `json:"updated_at" example:"2023-05-17T15:07:16.504+07:00"`
}

type TrainCarriageSeatResponses struct {
	TrainCarriageID uint                         `json:"train_carriage_id" example:"1"`
	Train           TrainResponse2               `json:"train"`
	Name            string                       `json:"name" example:"Gerbong 1"`
	Seat            []TrainSeatAvailableResponse `json:"seat"`
	CreatedAt       time.Time                    `json:"created_at" example:"2023-05-17T15:07:16.504+07:00"`
	UpdatedAt       time.Time                    `json:"updated_at" example:"2023-05-17T15:07:16.504+07:00"`
}

type TrainCarriageResponses struct {
	TrainCarriageID uint                `json:"train_carriage_id" example:"1"`
	Class           string              `json:"class" example:"Ekonomi"`
	Name            string              `json:"name" example:"Gerbong 1"`
	Seat            []TrainSeatResponse `json:"seat"`
}

type TrainCarriageResponseSimply struct {
	TrainCarriageID uint          `json:"train_carriage_id" example:"1"`
	Train           TrainResponse `json:"train"`
	Name            string        `json:"name" example:"Gerbong 1"`
}
