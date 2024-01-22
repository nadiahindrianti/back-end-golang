package models

import (
	"gorm.io/gorm"
)

type TravelerDetail struct {
    gorm.Model
    UserID        uint        `form:"user_id" json:"user_id"`
    User          User        `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
    TicketOrderID *uint       `gorm:"null" form:"ticket_order_id" json:"ticket_order_id"`
    TicketOrder   TicketOrder `gorm:"foreignKey:TicketOrderID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
    HotelOrderID  *uint       `gorm:"null" form:"hotel_order_id" json:"hotel_order_id"`
    HotelOrder    HotelOrder  `gorm:"foreignKey:HotelOrderID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
    Title         string      `form:"title" json:"title"`
    FullName      string      `form:"full_name" json:"full_name"`
    IDCardNumber  *string     `gorm:"null" form:"id_card_number" json:"id_card_number"`
}
