package main

import (
	"log"
	"workWithDB/pkg/api"
	"workWithDB/pkg/repository"

	"github.com/gorilla/mux"
)

func main() {
	db, err := repository.New(repository.ConnStr)
	if err != nil {
		log.Fatal(err.Error())
	}
	api := api.New(mux.NewRouter(), db)
	api.Handle()
	log.Fatal(api.ListenAndServe("localhost:8080"))
}
