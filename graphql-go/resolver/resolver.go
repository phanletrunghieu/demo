package resolver

import (
	"github.com/phanletrunghieu/demo/graphql-go/resolver/mutation"
	"github.com/phanletrunghieu/demo/graphql-go/resolver/query"
	"github.com/phanletrunghieu/demo/graphql-go/service"
)

// Resolver .
type Resolver struct {
	query.Query
	mutation.Mutation
}

// NewResolver .
func NewResolver(s *service.Service) *Resolver {
	return &Resolver{
		Query:    query.Query{Service: s},
		Mutation: mutation.Mutation{Service: s},
	}
}
