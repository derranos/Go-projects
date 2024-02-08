package api

import (
	"encoding/json"
	"infotecs/pkg/models"
	"net/http"

	"github.com/gorilla/mux"
)

func (api *api) create(w http.ResponseWriter, r *http.Request) {
	wallet, err := api.db.CreateWallet()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(wallet)
	if err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		return
	}
}

func (api *api) get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	walletIdStr, ok := vars["walletId"]
	if !ok {
		http.Error(w, "Invalid or missing walletId", 404)
		return
	}
	wallet, err := api.db.GetWallet(walletIdStr)
	if err != nil {
		http.Error(w, err.Error(), 404)
		return
	}
	err = json.NewEncoder(w).Encode(wallet)
	if err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		return
	}
}
func (api *api) send(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	walletIdStr, ok := vars["walletId"]
	if !ok {
		http.Error(w, "Invalid or missing walletId", 404)
		return
	}
	transfer := models.Transfer{}
	err := json.NewDecoder(r.Body).Decode(&transfer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	wallet, err := api.db.GetWallet(walletIdStr)
	if err != nil {
		http.Error(w, err.Error(), 404)
		return
	}
	if wallet.Balance >= transfer.Amount {
		_, err = api.db.GetWallet(transfer.TO)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		err = api.db.SendMoney(wallet.ID, transfer.TO, transfer.Amount)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = api.db.SaveTransfer(wallet.ID, transfer.TO, transfer.Amount)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	http.Error(w, "not enough money", 400)
}
func (api *api) getHistory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	walletIdStr, ok := vars["walletId"]
	if !ok {
		http.Error(w, "Invalid or missing walletId", 404)
		return
	}
	wallet, err := api.db.GetWallet(walletIdStr)
	if err != nil {
		http.Error(w, err.Error(), 404)
		return
	}
	data, err := api.db.GetHistory(wallet.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
