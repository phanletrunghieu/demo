package link

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

// Create implement Create for Link service
func (s *pgService) Create(l *domain.Link) error {
	return s.db.Create(l).Error
}

// Find implement Find for Link service
func (s *pgService) Find(id domain.UUID) (*domain.Link, error) {
	res := &domain.Link{
		Model: domain.Model{ID: id},
	}
	if err := s.db.Find(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}

// FindAll implement FindAll for Link service
func (s *pgService) FindAll(offset int32, limit int32) ([]*domain.Link, error) {
	res := []*domain.Link{}
	return res, s.db.Offset(offset).Limit(limit).Find(&res).Error
}
