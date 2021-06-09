package models

import (
	"fmt"
	"github.com/ehsaniara/go-crash/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/ehsaniara/go-crash/pkg/log"
	"time"
)

var db *gorm.DB

type Model struct {
	ID         int       `gorm:"primary_key" json:"id"`
	CreatedOn  time.Time `json:"createdOn" sql:"DEFAULT:'current_timestamp'"`
	ModifiedOn time.Time `json:"modifiedOn"`
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
		log.Log.Fatalf("db Error: %s", err)
	}

	if err := db.AutoMigrate(&Auth{}, &Customer{}); err != nil {
		log.Log.Fatalf("db AutoMigrate Error: %s\n", err)
		return
	}

	log.Log.Infof("Postgress DB Connection has started on Host: %s", config.AppConfig.DataBase.Host)
}
