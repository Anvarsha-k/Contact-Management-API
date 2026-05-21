package dto

// CreateContactRequest represents create contact payload.
type CreateContactRequest struct {
	FirstName  string `json:"first_name" validate:"required,min=2,max=100"`
	LastName   string `json:"last_name" validate:"max=100"`
	Email      string `json:"email" validate:"required,email"`
	PhoneNumber string `json:"phone_number" validate:"max=20"`
	CompanyName string `json:"company_name" validate:"max=255"`
	Status      string `json:"status" validate:"required,oneof=Active Inactive"`
}

// UpdateContactRequest represents update contact payload.
type UpdateContactRequest struct {
	FirstName   string `json:"first_name" validate:"omitempty,min=2,max=100"`
	LastName    string `json:"last_name" validate:"omitempty,max=100"`
	Email       string `json:"email" validate:"omitempty,email"`
	PhoneNumber string `json:"phone_number" validate:"omitempty,max=20"`
	CompanyName string `json:"company_name" validate:"omitempty,max=255"`
	Status      string `json:"status" validate:"omitempty,oneof=Active Inactive"`
}