package entity

import (
	"time"

	"gorm.io/gorm"
)

// Contact Model

type Contact struct {
	ID uint `gorm:"primaryKey" json:"id"`

	FirstName string `gorm:"type:varchar(100);not null" json:"first_name"`
	LastName  string `gorm:"type:varchar(100)" json:"last_name"`

	Email string `gorm:"type:varchar(255);not null" json:"email"`

	PhoneNumber string `gorm:"type:varchar(20)" json:"phone_number"`
	CompanyName string `gorm:"type:varchar(255)" json:"company_name"`

	Status string `gorm:"type:varchar(20);default:Active" json:"status"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}