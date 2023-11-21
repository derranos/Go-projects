package main

import (
	"log"
	"net/http"
	"testMod/figureServer/cmd/api"
)

func main() {
	api := api.New("localhost:8090", http.NewServeMux())
	api.FillEndpoints()
	log.Fatal(api.ListenAndServe())
}
