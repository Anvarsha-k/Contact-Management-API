package config

import (
	"fmt"
	"log"

	"github.com/Anvarsha-k/Contact-Management-API/internal/contact/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// for initialising Posgres database Connection
func ConnectDatabase(cfg *Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
		cfg.DBSSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}

	err = db.AutoMigrate(&entity.Contact{})
	if err != nil {
		return nil, fmt.Errorf("failed to migrate database: %w", err)
	}

	log.Println("database connection established")

	return db, nil
}
