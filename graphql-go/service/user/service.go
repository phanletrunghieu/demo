package user

import "github.com/phanletrunghieu/demo/graphql-go/domain"

// Service .
type Service interface {
	Create(u *domain.User) error
	Find(id domain.UUID) (*domain.User, error)
	FindByEmail(email string) (*domain.User, error)
	FindAll(offset int32, limit int32) ([]*domain.User, error)
}
