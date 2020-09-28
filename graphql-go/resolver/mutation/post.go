package mutation

import (
	"context"

	"github.com/phanletrunghieu/demo/graphql-go/domain"
	"github.com/phanletrunghieu/demo/graphql-go/resolver/model"
)

type postQueryArgs struct {
	URL         string
	Description string
}

// Post .
func (m *Mutation) Post(ctx context.Context, args postQueryArgs) (*model.LinkResolver, error) {
	link := &domain.Link{
		URL:         args.URL,
		Description: args.Description,
	}

	err := m.Service.LinkService.Create(link)

	return model.NewLinkResolver(ctx, m.Service, link), err
}
