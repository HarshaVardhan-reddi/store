package router

import (
	// "fmt"
	// "net/http"
	controller "store/controller/v1"

	"github.com/gorilla/mux"
)

func MainRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/hello",controller.ListStores)
	return r
}

// func helloRoute(w http.ResponseWriter, r *http.Request){
// 	fmt.Println("Hello route is here")
// }