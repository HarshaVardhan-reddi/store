package main

import (
	"log"
	"store/config"
	"store/router"
)

func main(){
	r := router.MainRouter()
	httpServer := config.ServerConfig.ConstructHttpServer(r) // constructing http server
	config.ConfigureMySQLServer()
	if err := httpServer.ListenAndServe(); err != nil{
		log.Fatal(err)
	}
}