package models

import (
	"time"

	"gorm.io/gorm"
)

type Rating struct {
	gorm.Model
	ID        uint       `gorm:"primaryKey;column:id" json:"id"`
	Rating    uint       `column:rating" json:"rating"`
	HotelID   uint       `column:hotel_id" json:"hotel_id"`
	UserID    uint       `column:user_id" json:"user_id"`
	Status    bool       `gorm:"default:true" json:"status"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
}
