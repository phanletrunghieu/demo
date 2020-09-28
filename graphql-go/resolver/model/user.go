package model

import (
	"context"
	"fmt"

	graphql "github.com/graph-gophers/graphql-go"

	"github.com/phanletrunghieu/demo/graphql-go/domain"
	"github.com/phanletrunghieu/demo/graphql-go/service"
)

// UserResolver .
type UserResolver struct {
	service *service.Service
	user    *domain.User
}

// NewUserResolver .
func NewUserResolver(ctx context.Context, service *service.Service, user *domain.User) *UserResolver {
	return &UserResolver{
		service: service,
		user:    user,
	}
}

// NewUsersResolver .
func NewUsersResolver(ctx context.Context, service *service.Service, users []*domain.User) []*UserResolver {
	var (
		ls = make([]*UserResolver, 0, len(users))
	)
	for _, user := range users {
		l := NewUserResolver(ctx, service, user)
		ls = append(ls, l)
	}
	return ls
}

// ID .
func (u *UserResolver) ID() graphql.ID {
	return graphql.ID(fmt.Sprint(u.user.ID))
}

// Name .
func (u *UserResolver) Name() string {
	return u.user.Name
}

// Email .
func (u *UserResolver) Email() string {
	return u.user.Email
}

// Links .
func (u *UserResolver) Links(ctx context.Context) []*LinkResolver {
	return NewLinksResolver(ctx, u.service, u.user.Links)
}
