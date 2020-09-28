package link

import "github.com/phanletrunghieu/demo/graphql-go/domain"

// Service .
type Service interface {
	Create(u *domain.Link) error
	Find(id domain.UUID) (*domain.Link, error)
	FindAll(offset int32, limit int32) ([]*domain.Link, error)
}
