package api

import (
	"infotecs/pkg/repository"
	"net/http"

	"github.com/gorilla/mux"
)

type api struct {
	r  *mux.Router
	db *repository.PGRepo
}

func New(router *mux.Router, db *repository.PGRepo) *api {
	return &api{r: router, db: db}
}

func (api *api) Handle() {
	api.r.HandleFunc("/api/v1/wallet", api.create).Methods("POST")
	api.r.HandleFunc("/api/v1/wallet/{walletId}/send", api.send).Methods("POST")
	api.r.HandleFunc("/api/v1/wallet/{walletId}", api.get).Methods("GET")
	api.r.HandleFunc("/api/v1/wallet/{walletId}/history", api.getHistory).Methods("GET")
}

func (api *api) ListenAndServe(addr string) error {
	return http.ListenAndServe(addr, api.r)
}
