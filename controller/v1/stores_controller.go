package v1

import (
	// "fmt"
	"encoding/json"
	"net/http"
	// "store/model"
	"store/services"
)

func ListStores(w http.ResponseWriter, r *http.Request){
	store_service := &services.StoreService{}
	stores := store_service.ListStores()
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(stores)
}