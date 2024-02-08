package main

import (
	"infotecs/pkg/api"
	"infotecs/pkg/repository"
	"log"

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
