package repository

import (
	"context"

	"github.com/Anvarsha-k/Contact-Management-API/internal/contact/entity"
)

// ContactRepository defines contact database operations.
type ContactRepository interface {

	Create(ctx context.Context, contact *entity.Contact) error
	
	FindActiveByEmail(ctx context.Context,email string,) (*entity.Contact, error)
}