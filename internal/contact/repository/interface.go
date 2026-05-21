package repository

import (
	"context"

	"github.com/Anvarsha-k/Contact-Management-API/internal/contact/dto"
	"github.com/Anvarsha-k/Contact-Management-API/internal/contact/entity"
)

// ContactRepository defines contact database operations

type ContactRepository interface {
	Create(ctx context.Context, contact *entity.Contact) error

	FindActiveByEmail(ctx context.Context, email string) (*entity.Contact, error)

	// fetches paginated contacts
	List(ctx context.Context, query dto.ContactListQuery) ([]entity.Contact, int64, error)

	// fetches contact by ID.
	GetByID(ctx context.Context, id uint,) (*entity.Contact, error)

	Update(ctx context.Context, contact *entity.Contact) error

	//soft delete.
	Delete(ctx context.Context, contact *entity.Contact) error
}
