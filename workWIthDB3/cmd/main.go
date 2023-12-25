package main

import (
	"log"
	"log/slog"
	"os"
	"workWithDB/pkg/api"
	"workWithDB/pkg/repository"

	"github.com/gorilla/mux"
)

func main() {
	db, err := repository.New(repository.ConnStr)
	if err != nil {
		log.Fatal(err.Error())
	}
	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))
	api := api.New(mux.NewRouter(), db, logger)
	api.Handle()
	log.Fatal(api.ListenAndServe("localhost:8080"))
}
