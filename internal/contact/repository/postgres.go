package repository

import (
	"context"
	"errors"

	"github.com/Anvarsha-k/Contact-Management-API/internal/contact/entity"
	"gorm.io/gorm"
)


type contactRepository struct {
	db *gorm.DB
}

func NewContactRepository(db *gorm.DB) ContactRepository {
	return &contactRepository{
		db: db,
	}
}

// inserts new contact record.

func (r *contactRepository) Create(ctx context.Context, contact *entity.Contact,) error {

	return r.db.WithContext(ctx).Create(contact).Error
}

// checking existing active contact by email.
func (r *contactRepository) FindActiveByEmail(ctx context.Context, email string,) (*entity.Contact, error) {

	var contact entity.Contact

	err := r.db.WithContext(ctx).Where("email = ? AND status = ?", email, "Active").First(&contact).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &contact, nil
}