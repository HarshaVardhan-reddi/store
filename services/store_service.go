package services

import (
	"fmt"
	"log"
	"store/config"
	"store/model"

	// "gorm.io/gorm"
)

type StoreService struct{
}

func (s *StoreService) ListStores() *[]model.Store {
	var (
		stores []model.Store;
		res []model.Store
	)
	result := config.DbConn.Find(&stores)
	rows, err := result.Rows()
	if err != nil{
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next(){
		var store model.Store
    if err := config.DbConn.ScanRows(rows, &store); err != nil {
			log.Fatal(err)
    }
		res = append(res, store)
	}
	for index, val := range(res){
		fmt.Println("index:",index, "val:",val)
	}
	return &res
}

func (s *StoreService) AddStore(store *model.Store) (*model.Store, error){
	result := config.DbConn.Create(&store)
	if err := result.Error; err != nil{
		return &model.Store{}, err
	}
	return store, nil
}


func (s *StoreService) DummyStore() *model.Store {
	return &model.Store{}
}