package config

import (
	"log"
	"os"

	"github.com/goccy/go-yaml"
)

type DatabaseServer struct{
	Host string `yaml:"host"`
  Port int64 `yaml:"port"`
  Database string `yaml:"database"`
  User string `yaml:"user"`
  Password string `yaml:"password"`
}

type DatabaseConfig struct{
	Development DatabaseServer
	Test DatabaseServer
	Production DatabaseServer
}

func init(){
	content, err := os.ReadFile("config/database.yml")
	if err != nil {
		log.Fatal("Database configuration",err)
	}
	db_config := DatabaseConfig{}
	if err := yaml.Unmarshal(content,&db_config); err != nil{
		log.Fatal("Datatabase configuration",err)
	}
}