package v1

import (
	// "fmt"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	// "store/model"
	"store/services"

	"github.com/gorilla/mux"
)

func ListStores(w http.ResponseWriter, r *http.Request){
	store_service := &services.StoreService{}
	stores := store_service.ListStores()
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(stores)
}

func CreateStore(w http.ResponseWriter, r *http.Request){
	w.Header().Add("Content-Type","application/json")
	var response map[string]string = make(map[string]string, 0)
	storeService := services.StoreService{}
	store := storeService.DummyStore()

	if err := json.NewDecoder(r.Body).Decode(&store); err != nil{
		log.Fatal(err)
	}

	defer r.Body.Close()

	if _, err := storeService.AddStore(store); err != nil{
		w.WriteHeader(http.StatusPreconditionFailed)
		json.NewEncoder(w).Encode(map[string]string{"message":err.Error()})
		return 
	}
	

	response["message"] = "Successfully created the store"
	w.WriteHeader(http.StatusOK)

	log.Println(response)
	json.NewEncoder(w).Encode(response)
}

func GetStore(w http.ResponseWriter, r *http.Request){
	w.Header().Add("Content-Type", "application/json")
	pathvars := mux.Vars(r)
	id, err := strconv.Atoi(pathvars["id"])
	if err != nil{
		log.Fatal(err)
	}
	store_service := services.StoreService{}
	store, errFinding := store_service.FindStoreWithId(int64(id))
	if errFinding != nil{
		w.WriteHeader(http.StatusPreconditionFailed)
		json.NewEncoder(w).Encode(map[string]string{"message":err.Error()})
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(store)
}