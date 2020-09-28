package http

import (
	"net/http"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/phanletrunghieu/demo/graphql-go/auth"
)

// NewGraphQLHandler .
func NewGraphQLHandler(graphqlSchema string, resolver interface{}) http.Handler {
	schema := graphql.MustParseSchema(graphqlSchema, resolver)
	return auth.MiddlewareGraphQL(&relay.Handler{Schema: schema})
}
