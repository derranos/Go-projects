package api

import (
	"log/slog"
	"net/http"
	"workWithDB/pkg/repository"

	"github.com/gorilla/mux"
)

type api struct {
	r      *mux.Router
	db     *repository.PGRepo
	logger *slog.Logger
}

func New(router *mux.Router, db *repository.PGRepo, logger *slog.Logger) *api {
	return &api{r: router, db: db, logger: logger}
}

func (api *api) Handle() {
	api.r.HandleFunc("/api/books", api.books).Queries("id", "{id}")
	api.r.HandleFunc("/api/books", api.books)
	api.r.HandleFunc("/api/authors", api.authors).Queries("id", "{id}")
	api.r.HandleFunc("/api/authors", api.authors)
	api.r.HandleFunc("/api/genres", api.genres).Queries("id", "{id}")
	api.r.HandleFunc("/api/genres", api.genres)
	api.r.Use(api.middleware)
}

func (api *api) ListenAndServe(addr string) error {
	return http.ListenAndServe(addr, api.r)
}
