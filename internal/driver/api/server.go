package api

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/samber/do"

	"matcha/internal/driver/api/gqlgen"
)

func NewServer(i *do.Injector) (*http.Server, error) {
	resolver := do.MustInvoke[*Resolver](i)
	r := chi.NewRouter()

	r.Get("/", playground.Handler("GraphQL", "/query").ServeHTTP)
	r.Post("/query", handler.NewDefaultServer(gqlgen.NewExecutableSchema(gqlgen.Config{Resolvers: resolver})).ServeHTTP)

	return &http.Server{
		Handler: r,
		Addr:    do.MustInvokeNamed[string](i, "ServerAddress"),
	}, nil
}
