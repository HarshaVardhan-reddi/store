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
	return r
}

// func helloRoute(w http.ResponseWriter, r *http.Request){
// 	fmt.Println("Hello route is here")
// }