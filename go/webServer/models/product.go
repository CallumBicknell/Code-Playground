package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Product struct {
	Product_ID   uuid.UUID `gorm:"type:uuid;primary_key;product_id;"`
	Lister       User.ID
	Product_Name string
	Price        uint64
	Rating       uint8
	Image        string
	Active       bool      `gorm:"default:true;active;"`
	UpdatedAt    time.Time `gorm:"updated_at;"`
	CreatedAt    time.Time `gorm:"updated_at;"`
}
