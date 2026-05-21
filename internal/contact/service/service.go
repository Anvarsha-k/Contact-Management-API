package service

import (
	"context"
	"errors"
	"math"
	"strings"

	"github.com/Anvarsha-k/Contact-Management-API/internal/contact/dto"
	"github.com/Anvarsha-k/Contact-Management-API/internal/contact/entity"
	"github.com/Anvarsha-k/Contact-Management-API/internal/contact/repository"
	"github.com/Anvarsha-k/Contact-Management-API/internal/shared/pagination"
	"github.com/gofiber/fiber/v2"
)

type contactService struct {
	repository repository.ContactRepository
}

func NewContactService(repository repository.ContactRepository) ContactService {

	return &contactService{
		repository: repository,
	}
}

// CreateContact handles contact creation business logic

func (s *contactService) CreateContact(ctx context.Context, request dto.CreateContactRequest) (*dto.ContactResponse, error) {

	normalizedEmail := strings.TrimSpace(
		strings.ToLower(request.Email),
	)

	existingContact, err := s.repository.FindActiveByEmail(ctx, normalizedEmail)

	if err != nil {
		return nil, err
	}

	if existingContact != nil {
		return nil, errors.New("active contact with this email already exists")
	}

	contact := entity.Contact{

		FirstName:   strings.TrimSpace(request.FirstName),
		LastName:    strings.TrimSpace(request.LastName),
		Email:       strings.TrimSpace(strings.ToLower(request.Email)),
		PhoneNumber: strings.TrimSpace(request.PhoneNumber),
		CompanyName: strings.TrimSpace(request.CompanyName),
		Status:      request.Status,
	}

	err = s.repository.Create(ctx, &contact)
	if err != nil {
		return nil, err
	}

	response := mapToContactResponse(contact)

	return &response, nil
}

// ListContacts fetches paginated contacts
func (s *contactService) ListContacts(ctx context.Context, query dto.ContactListQuery) (interface{}, error) {

	contacts, totalRows, err := s.repository.List(ctx, query)

	if err != nil {
		return nil, err
	}

	var response []dto.ContactResponse

	for _, contact := range contacts {

		response = append(response, mapToContactResponse(contact))
	}

	totalPages := int(math.Ceil(
		float64(totalRows) / float64(query.Limit),
	))

	return fiber.Map{
		"contacts": response,
		"pagination": pagination.Pagination{
			Page:       query.Page,
			Limit:      query.Limit,
			TotalRows:  totalRows,
			TotalPages: totalPages,
		},
	}, nil
}

// GetContactByID fetches single contact
func (s *contactService) GetContactByID(ctx context.Context, id uint) (*dto.ContactResponse, error) {

	contact, err := s.repository.GetByID(ctx, id)

	if err != nil {
		return nil, err
	}

	if contact == nil {
		return nil, errors.New("contact not found")
	}

	response := mapToContactResponse(*contact)

	return &response, nil
}

// Update Contactr

func (s *contactService) UpdateContact(ctx context.Context, id uint, request dto.UpdateContactRequest) (*dto.ContactResponse, error) {

	contact, err := s.repository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if contact == nil {
		return nil, errors.New("contact not found")
	}

	// Normalise email
	if request.Email != "" {

		normalizedEmail := strings.TrimSpace(strings.ToLower(request.Email))

		existingContact, err := s.repository.FindActiveByEmail(ctx, normalizedEmail)

		if err != nil {
			return nil, err
		}

		if existingContact != nil &&
			existingContact.ID != contact.ID {

			return nil, errors.New(
				"active contact with this email already exists",
			)
		}

		contact.Email = normalizedEmail
	}

	if request.FirstName != "" {
		contact.FirstName = strings.TrimSpace(request.FirstName)
	}

	if request.LastName != "" {
		contact.LastName = strings.TrimSpace(request.LastName)
	}

	if request.PhoneNumber != "" {
		contact.PhoneNumber = strings.TrimSpace(request.PhoneNumber)
	}

	if request.CompanyName != "" {
		contact.CompanyName = strings.TrimSpace(request.CompanyName)
	}

	if request.Status != "" {
		contact.Status = request.Status
	}

	err = s.repository.Update(ctx, contact)
	if err != nil {
		return nil, err
	}

	response := mapToContactResponse(*contact)

	return &response, nil
}

func (s *contactService) DeleteContact( ctx context.Context, id uint) error {

	contact, err := s.repository.GetByID( ctx,id)

	if err != nil {
		return err
	}

	if contact == nil {
		return errors.New("contact not found")
	}

	return s.repository.Delete(ctx, contact)
}