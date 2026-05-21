package service

import (
	"github.com/Anvarsha-k/Contact-Management-API/internal/contact/dto"
	"github.com/Anvarsha-k/Contact-Management-API/internal/contact/entity"
)

//converts entity to response DTO

func mapToContactResponse( contact entity.Contact) dto.ContactResponse {

	return dto.ContactResponse{
		ID: contact.ID,
		FirstName: contact.FirstName,
		LastName:  contact.LastName,
		Email: contact.Email,
		PhoneNumber: contact.PhoneNumber,
		CompanyName: contact.CompanyName,
		Status: contact.Status,
		CreatedAt: contact.CreatedAt,
		UpdatedAt: contact.UpdatedAt,
	}
}