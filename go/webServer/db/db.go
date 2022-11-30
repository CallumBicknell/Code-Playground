package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/CallumBicknell/go-webServer/config"
	"github.com/jinzhu/gorm"
)

var db *sql.DB

func Init() {
	config := config.GetConfig()
	host := config.GetString("db.host")
	port := config.GetInt("db.port")
	user := config.GetString("db.username")
	password := config.GetString("db.password")
	dbname := config.GetString("db.name")

	dbinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := gorm.Open("postgres", dbinfo)
	if err != nil {
		fmt.Fprintf(os.Stderr, "db connection failed: %s", err.Error())
		os.Exit(1)
	}

	db.LogMode(true)
	defer db.Close()

	// DropTables() // Only if developing otherwise have a backup!
	// CreateTables()
}

func GetDB() *sql.DB {
	return db
}
