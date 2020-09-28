package query

import (
	"context"

	"github.com/phanletrunghieu/demo/graphql-go/resolver/model"
)

type feedQueryArgs struct {
	Filter *string
	Offset *int32
	Limit  *int32
}

// Feed .
func (r *Query) Feed(ctx context.Context, filter feedQueryArgs) ([]*model.LinkResolver, error) {
	links, err := r.Service.LinkService.FindAll(*filter.Offset, *filter.Limit)
	if err != nil {
		return nil, err
	}

	linksResolver := model.NewLinksResolver(ctx, r.Service, links)

	return linksResolver, nil
}
