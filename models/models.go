package models

import (
	"fmt"
	"github.com/ehsaniara/go-crash/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
	DeletedOn  int `json:"deleted_on"`
}

// Setup initializes the database instance
func Setup() {

	var err error
	// https://github.com/go-gorm/postgres
	db, err = gorm.Open(postgres.New(postgres.Config{
		DSN: fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=UTC",
			config.AppConfig.DataBase.Host,
			config.AppConfig.DataBase.Username,
			config.AppConfig.DataBase.Password,
			config.AppConfig.DataBase.DbName,
			config.AppConfig.DataBase.Port),
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		log.Fatalf("db Error: %s\n", err)
	}

	if err := db.AutoMigrate(&Auth{}, &Customer{}); err != nil {
		log.Fatalf("db AutoMigrate Error: %s\n", err)
		return
	}

	fmt.Printf(" - Postgress DB Connection has started on Host: %s\n", config.AppConfig.DataBase.Host)
}
