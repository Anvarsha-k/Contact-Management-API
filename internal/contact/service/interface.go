package service

import (
	"context"

	"github.com/Anvarsha-k/Contact-Management-API/internal/contact/dto"
)

// ContactService defines contact business logic
type ContactService interface {

	CreateContact(ctx context.Context, request dto.CreateContactRequest,) (*dto.ContactResponse, error)

	ListContacts(ctx context.Context, query dto.ContactListQuery) (interface{}, error)
	
	GetContactByID( ctx context.Context, id uint) (*dto.ContactResponse, error)

	UpdateContact( ctx context.Context, id uint, request dto.UpdateContactRequest) (*dto.ContactResponse, error)

	DeleteContact( ctx context.Context, id uint) error
}