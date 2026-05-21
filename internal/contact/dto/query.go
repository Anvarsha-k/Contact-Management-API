package dto

// ContactListQuery represents contact listing query params

type ContactListQuery struct {
	Page   int
	Limit  int
	Search string
	Status string
	SortBy string
	Order  string
}