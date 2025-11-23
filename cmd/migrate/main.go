package main

import (
	"flag"
	"fmt"
	"log"
	"store/config"

	"github.com/pressly/goose/v3"
	_ "store/db/migrations"
)

func main(){
	config.ConfigureMySQLServer()
	db, err := config.DbConn.DB()

	goose.SetDialect("mysql")
	
	if err != nil{
		log.Fatal(err)
	}
	mode := flag.String("mode","","whats the mode of the goose? up or down")
	path := flag.String("path","","path for migration folder")

	flag.Parse()
	
	fmt.Println("mode:",*mode)
	fmt.Println("path:",*path)
	if *mode == "up" && *path != ""{
		if err := goose.Up(db,*path); err != nil{
			log.Fatal(err)
		}
	}else if *mode == "down" && *path != ""{
		if err := goose.Down(db,*path); err != nil{
			log.Fatal(err)
		}
	}else if *mode == "rollback" && *path != ""{
		if err := goose.Redo(db,*path); err != nil{
			log.Fatal(err)
		}
	}
}