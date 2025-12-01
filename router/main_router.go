package router

import (
	// "fmt"
	// "net/http"
	controller "store/controllers/v1"

	"github.com/gorilla/mux"
)

func MainRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/stores",controller.ListStores).Methods("GET")
	r.HandleFunc("/stores",controller.CreateStore).Methods("POST")
	r.HandleFunc("/stores/{id}",controller.GetStore).Methods("GET")
	r.HandleFunc("/stores/{id}",controller.UpdateStore).Methods("PATCH")
	return r
}

// func helloRoute(w http.ResponseWriter, r *http.Request){
// 	fmt.Println("Hello route is here")
// }