package model

import (
	"context"
	"fmt"

	graphql "github.com/graph-gophers/graphql-go"

	"github.com/phanletrunghieu/demo/graphql-go/domain"
	"github.com/phanletrunghieu/demo/graphql-go/service"
)

// LinkResolver .
type LinkResolver struct {
	service *service.Service
	link    *domain.Link
}

// NewLinkResolver .
func NewLinkResolver(ctx context.Context, service *service.Service, link *domain.Link) *LinkResolver {
	return &LinkResolver{
		service: service,
		link:    link,
	}
}

// NewLinksResolver .
func NewLinksResolver(ctx context.Context, service *service.Service, links []*domain.Link) []*LinkResolver {
	var (
		ls = make([]*LinkResolver, 0, len(links))
	)
	for _, link := range links {
		l := NewLinkResolver(ctx, service, link)
		ls = append(ls, l)
	}
	return ls
}

// ID .
func (l *LinkResolver) ID() graphql.ID {
	return graphql.ID(fmt.Sprint(l.link.ID))
}

// Description .
func (l *LinkResolver) Description() string {
	return l.link.Description
}

// URL .
func (l *LinkResolver) URL() string {
	return l.link.URL
}

// PostedBy .
func (l *LinkResolver) PostedBy(ctx context.Context) (*UserResolver, error) {
	user, err := l.service.UserService.Find(l.link.ID)
	if err != nil {
		return nil, err
	}
	return NewUserResolver(ctx, l.service, user), nil
}
