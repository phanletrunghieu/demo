package service

import (
	"github.com/jinzhu/gorm"
	"github.com/phanletrunghieu/demo/graphql-go/domain"
	"github.com/phanletrunghieu/demo/graphql-go/service/link"
	"github.com/phanletrunghieu/demo/graphql-go/service/user"
)

// Service .
type Service struct {
	UserService user.Service
	LinkService link.Service
}

// NewService .
func NewService(db *gorm.DB) (*Service, error) {
	if !db.HasTable(&domain.User{}) {
		if err := db.CreateTable(&domain.User{}).Error; err != nil {
			return nil, err
		}
	}

	if !db.HasTable(&domain.Link{}) {
		if err := db.CreateTable(&domain.Link{}).Error; err != nil {
			return nil, err
		}
		db.Model(&domain.Link{}).AddForeignKey("posted_by", "users(id)", "RESTRICT", "RESTRICT")
	}

	return &Service{
		UserService: user.NewPGService(db),
		LinkService: link.NewPGService(db),
	}, nil
}
