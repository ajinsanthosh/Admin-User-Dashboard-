package db

import (
	"example/pkg/config"
	"example/pkg/domain"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase(cfg *config.Config) (*gorm.DB, error) {

	psqlInfo := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s", cfg.DBUser, cfg.DBName, cfg.DBPassword, cfg.DBHost, cfg.DBPort)

	db, dberr := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if dberr != nil {
		fmt.Println("ERROR  in db connection", dberr)
	}
	db.AutoMigrate(&domain.Users{})
	db.AutoMigrate(&domain.Admin{})

	return db, nil
}
