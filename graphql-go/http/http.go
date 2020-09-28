package http

import (
	"net/http"

	"github.com/go-chi/chi"
)

// NewHTTPHandler .
func NewHTTPHandler(graphqlSchema string, resolver interface{}) http.Handler {
	r := chi.NewRouter()

	r.Handle("/graphql", NewGraphQLHandler(graphqlSchema, resolver))
	r.Handle("/graphiql", NewGraphiQLHandler())

	// file server
	fs := http.FileServer(http.Dir("./public"))
	r.Handle("/", fs)

	return r
}
