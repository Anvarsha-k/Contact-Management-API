package service

import (
	"context"

	"github.com/Anvarsha-k/Contact-Management-API/internal/contact/dto"
)

// ContactService defines contact business logic.
type ContactService interface {

	CreateContact(ctx context.Context, request dto.CreateContactRequest,) (*dto.ContactResponse, error)
}