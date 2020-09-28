package user

import (
	"github.com/jinzhu/gorm"
	"github.com/phanletrunghieu/demo/graphql-go/domain"
)

type pgService struct {
	db *gorm.DB
}

// NewPGService .
func NewPGService(db *gorm.DB) Service {
	return &pgService{
		db: db,
	}
}

// Create implement Create for User service
func (s *pgService) Create(u *domain.User) error {
	return s.db.Create(u).Error
}

// Find implement Find for User service
func (s *pgService) Find(id domain.UUID) (*domain.User, error) {
	res := &domain.User{
		Model: domain.Model{ID: id},
	}
	if err := s.db.Find(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}

// Find implement Find for User service
func (s *pgService) FindByEmail(email string) (*domain.User, error) {
	res := &domain.User{}
	if err := s.db.Where("email = ?", email).First(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}

// FindAll implement FindAll for User service
func (s *pgService) FindAll(offset int32, limit int32) ([]*domain.User, error) {
	res := []*domain.User{}
	return res, s.db.Offset(offset).Limit(limit).Find(&res).Error
}
