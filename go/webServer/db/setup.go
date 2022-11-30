package db

// TODO: Prevent this import cycle face🤦‍♂️

import (
	// _ "github.com/CallumBicknell/go-webServer/models"
	"github.com/jinzhu/gorm"
)

func CreateTables(db *gorm.DB) {
}

func DropTables(db *gorm.DB) {
	// db.DropTableIfExists(&models.User{})
	// db.DropTableIfExists()
}
