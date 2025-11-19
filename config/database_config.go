package config

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/goccy/go-yaml"
	ormsql "gorm.io/driver/mysql"
	"gorm.io/gorm"
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

var db_config = DatabaseConfig{}

func init(){
	content, err := os.ReadFile("config/database.yml")
	if err != nil {
		log.Fatal("Database configuration",err)
	}
	if err := yaml.Unmarshal(content,&db_config); err != nil{
		log.Fatal("Datatabase configuration",err)
	}
}

func ConfigureMySQLServer(){
	databaseServer, err := fetchEnvBasedDatabase(&db_config, "development")
	if err != nil {
		log.Fatal("Database configuration", err)
	}
	rawMysqlConfig := mysql.Config{
		User: databaseServer.User,
		Passwd: databaseServer.Password,
		DBName: databaseServer.Database,
		Addr: fmt.Sprintf("%s:%d",databaseServer.Host, databaseServer.Port),
		Timeout: 10 * time.Second,
		ReadTimeout: 10 * time.Second,
		ParseTime: true,
    AllowNativePasswords: true,
	}
	MySQLConfig := ormsql.Config{
		DSNConfig: &rawMysqlConfig,
	}
	db, errInDBCon := gorm.Open(ormsql.New(MySQLConfig), &gorm.Config{})
	if(errInDBCon != nil){
		log.Fatal("Database configuration", err)
	}
	log.Println("DB::",db)
}

func fetchEnvBasedDatabase(database_config *DatabaseConfig, env string) (*DatabaseServer, error) {
	switch(env){
	case "development":
		return &database_config.Development, nil
	default:
		return &DatabaseServer{}, errors.New("database configuration not found for the requesting environment")
	}
}