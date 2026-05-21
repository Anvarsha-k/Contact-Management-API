package service

import (
	"context"
	"errors"
	"strings"

	"github.com/Anvarsha-k/Contact-Management-API/internal/contact/dto"
	"github.com/Anvarsha-k/Contact-Management-API/internal/contact/entity"
	"github.com/Anvarsha-k/Contact-Management-API/internal/contact/repository"
)

type contactService struct {
	repository repository.ContactRepository
}


func NewContactService(repository repository.ContactRepository,) ContactService {

	return &contactService{
		repository: repository,
	}
}

// CreateContact handles contact creation business logic.

func (s *contactService) CreateContact(ctx context.Context, request dto.CreateContactRequest,) (*dto.ContactResponse, error) {

	existingContact, err := s.repository.FindActiveByEmail(ctx, request.Email,)

	if err != nil {
		return nil, err
	}

	if existingContact != nil {
		return nil, errors.New("active contact with this email already exists")
	}

	contact := entity.Contact{

		FirstName: strings.TrimSpace(request.FirstName),
		LastName: strings.TrimSpace(request.LastName),
		Email: strings.TrimSpace(strings.ToLower(request.Email)),
		PhoneNumber: strings.TrimSpace(request.PhoneNumber),
		CompanyName: strings.TrimSpace(request.CompanyName),
		Status: request.Status,
	}

	err = s.repository.Create(ctx, &contact)
	if err != nil {
		return nil, err
	}

	response := &dto.ContactResponse{

		ID: contact.ID,
		FirstName: contact.FirstName,
		LastName: contact.LastName,
		Email: contact.Email,
		PhoneNumber: contact.PhoneNumber,
		CompanyName: contact.CompanyName,
		Status: contact.Status,
		CreatedAt: contact.CreatedAt,
		UpdatedAt: contact.UpdatedAt,
	}

	return response, nil
}