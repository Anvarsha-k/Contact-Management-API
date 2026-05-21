package dto

import "time"

// ContactResponse represents API response payload.
type ContactResponse struct {
	ID uint `json:"id"`

	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`

	Email string `json:"email"`

	PhoneNumber string `json:"phone_number"`
	CompanyName string `json:"company_name"`

	Status string `json:"status"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
