package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/samber/do"
)

func NewServer(i *do.Injector) (*http.Server, error) {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})

	return &http.Server{
		Handler: r,
		Addr:    do.MustInvokeNamed[string](i, "ServerAddress"),
	}, nil
}
