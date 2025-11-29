package model

import (
	"errors"

	"gorm.io/gorm"
)

// import "gorm.io/gorm"

type Store struct{
	// gorm.Model
	Id int64
	Name string
	Description string
	TagLine string
	StoreCode string `gorm:"column:store_code"`
}

func (s *Store) isIdAlreadyPresent() (bool, error){
	if s.Id > 0{
		return false, errors.New("ID cannot be inserted into db")
	}
	return true, nil
}

func (s *Store) isMandatoryAttributesExists() (bool, error){
	if s.Name == ""{
		return false, errors.New("name is mandatory for creating a new store")
	}
	if s.Description == ""{
		return false, errors.New("description is mandatory")
	}
	return true, nil
}


func (s *Store) BeforeCreate(tx *gorm.DB) (err error){
	if ok, err := s.isIdAlreadyPresent(); !ok{
		return err
	}
	if ok, err := s.isMandatoryAttributesExists(); !ok{
		return err
	}
	return 
}