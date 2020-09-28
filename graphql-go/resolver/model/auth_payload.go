package model

import (
	"context"

	"github.com/phanletrunghieu/demo/graphql-go/domain"
	"github.com/phanletrunghieu/demo/graphql-go/service"
)

// AuthPayloadResolver .
type AuthPayloadResolver struct {
	service     *service.Service
	authPayload *domain.AuthPayload
}

// NewAuthPayloadResolver .
func NewAuthPayloadResolver(ctx context.Context, service *service.Service, authPayload *domain.AuthPayload) *AuthPayloadResolver {
	return &AuthPayloadResolver{
		service:     service,
		authPayload: authPayload,
	}
}

// NewAuthPayloadsResolver .
func NewAuthPayloadsResolver(ctx context.Context, service *service.Service, authPayloads []*domain.AuthPayload) []*AuthPayloadResolver {
	var (
		ls = make([]*AuthPayloadResolver, 0, len(authPayloads))
	)
	for _, link := range authPayloads {
		l := NewAuthPayloadResolver(ctx, service, link)
		ls = append(ls, l)
	}
	return ls
}

// Token .
func (l *AuthPayloadResolver) Token() string {
	return l.authPayload.Token
}

// User .
func (l *AuthPayloadResolver) User(ctx context.Context) *UserResolver {
	return NewUserResolver(ctx, l.service, l.authPayload.User)
}
