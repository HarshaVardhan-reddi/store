package main

import (
	"fmt"
	"log"
	"net/http"
	"store/config"
	"github.com/gorilla/mux"
)

func main(){
	//TODO:Need to move into a different conf, basically under routes folder 
	r := mux.NewRouter() 
	r.HandleFunc("/hello", helloForTesting).Methods("GET")
	// fmt.Println() config.DatabaseConfig

	// Http server conf construction goes here
	httpServer := config.ServerConfig.ConstructHttpServer(r) // constructing http server
	config.ConfigureMySQLServer()
	if err := httpServer.ListenAndServe(); err != nil{
		log.Fatal(err)
	}
}

//TODO:Need to remove this function, added for testing 
func helloForTesting(w http.ResponseWriter, r *http.Request){
	fmt.Println("hello from /hello")
}