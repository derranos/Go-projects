package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"workWithDB/pkg/models"

	"github.com/gorilla/mux"
)

func (api *api) genres(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		vars := mux.Vars(r)
		stringID, ok := vars["id"]
		if ok {
			id, err := strconv.Atoi(stringID)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			data, err := api.db.GetGenreByID(id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			err = json.NewEncoder(w).Encode(data)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			return
		}
		data, err := api.db.GetGenres()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = json.NewEncoder(w).Encode(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	case http.MethodPost:
		var genre models.Genre
		err := json.NewDecoder(r.Body).Decode(&genre)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = api.db.CreateGenre(genre)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	case http.MethodDelete:
		vars := mux.Vars(r)
		stringID, ok := vars["id"]
		if ok {
			id, err := strconv.Atoi(stringID)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			err = api.db.DeleteGenreByID(id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			return
		}
	}
}
