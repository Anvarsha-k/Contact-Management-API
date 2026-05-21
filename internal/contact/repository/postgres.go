package repository

import (
	"context"
	"strings"

	"github.com/Anvarsha-k/Contact-Management-API/internal/contact/dto"
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

// inserts new contact record

func (r *contactRepository) Create(ctx context.Context, contact *entity.Contact) error {

	return r.db.WithContext(ctx).Create(contact).Error
}

// checking existing active contact by email
func (r *contactRepository) FindActiveByEmail(ctx context.Context, email string) (*entity.Contact, error) {

	var contact entity.Contact

	result := r.db.WithContext(ctx).Where("LOWER(email) = ? AND status = ?", strings.ToLower(email), "Active").First(&contact)

	if result.RowsAffected == 0 {
		return nil, nil
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return &contact, nil
}

// List fetches paginated contacts
func (r *contactRepository) List(ctx context.Context, query dto.ContactListQuery) ([]entity.Contact, int64, error) {

	var contacts []entity.Contact
	var totalRows int64

	db := r.db.WithContext(ctx).
		Model(&entity.Contact{})

	if query.Search != "" {

		searchPattern := "%" + query.Search + "%"

		db = db.Where(
			`first_name ILIKE ?
			OR last_name ILIKE ?
			OR email ILIKE ?`,
			searchPattern,
			searchPattern,
			searchPattern,
		)
	}

	// Filter by status
	if query.Status != "" {
		db = db.Where("status = ?", query.Status)
	}

	// Count total records
	err := db.Count(&totalRows).Error
	if err != nil {
		return nil, 0, err
	}

	// sorting
	allowedSortFields := map[string]bool{
		"created_at": true,
		"first_name": true,
		"email":      true,
		"status":     true,
	}

	sortBy := "created_at"

	if allowedSortFields[query.SortBy] {
		sortBy = query.SortBy
	}

	order := "desc"

	if query.Order == "asc" {
		order = "asc"
	}

	offset := (query.Page - 1) * query.Limit

	err = db.Order(sortBy + " " + order).Limit(query.Limit).Offset(offset).Find(&contacts).Error

	if err != nil {
		return nil, 0, err
	}

	return contacts, totalRows, nil
}

// GetByID fetches contact by ID
func (r *contactRepository) GetByID(ctx context.Context, id uint) (*entity.Contact, error) {

	var contact entity.Contact

	result := r.db.WithContext(ctx).First(&contact, id)

	if result.RowsAffected == 0 {
		return nil, nil
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return &contact, nil
}

// Update Contacts
func (r *contactRepository) Update(ctx context.Context, contact *entity.Contact) error {

	return r.db.WithContext(ctx).Save(contact).Error
}

// Delete performs soft delete
func (r *contactRepository) Delete( ctx context.Context, contact *entity.Contact) error {

	return r.db.WithContext(ctx).Delete(contact).Error
}