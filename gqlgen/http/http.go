package http

import (
	"net/http"

	"github.com/99designs/gqlgen/handler"
	"github.com/go-chi/chi"

	prisma "github.com/phanletrunghieu/demo/gqlgen/prisma-client"
	"github.com/phanletrunghieu/demo/gqlgen/resolver"
)

// NewHTTPHandler .
func NewHTTPHandler() http.Handler {
	router := chi.NewRouter()

	client := prisma.New(nil)
	r := resolver.Resolver{
		Prisma: client,
	}
	router.Handle("/graphql", handler.GraphQL(resolver.NewExecutableSchema(resolver.Config{Resolvers: &r})))
	router.Handle("/graphiql", handler.Playground("GraphQL Playground", "/graphql"))

	// file server
	// fs := http.FileServer(http.Dir("./public"))
	// r.Handle("/", fs)

	return router
}
