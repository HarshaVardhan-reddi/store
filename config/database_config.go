package config

import (
	"database/sql"
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

var DbConfig = DatabaseConfig{}
var Database *sql.DB
var DbConn *gorm.DB

func init(){
	content, err := os.ReadFile("config/database.yml")
	if err != nil {
		log.Fatal("Database configuration",err)
	}
	if err := yaml.Unmarshal(content,&DbConfig); err != nil{
		log.Fatal("Datatabase configuration",err)
	}
}

func ConfigureMySQLServer(){
	databaseServer, err := fetchEnvBasedDatabase("development")
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
	var errInDBCon error
	DbConn, errInDBCon = gorm.Open(ormsql.New(MySQLConfig), &gorm.Config{})
	if(errInDBCon != nil){
		log.Fatal("Database configuration", err)
	}

	var dbErr error
	Database, dbErr = DbConn.DB()
	if dbErr != nil {
			log.Fatal("Failed to get DB instance:", err)
	}
	if pingErr := Database.Ping(); pingErr != nil{
		log.Fatal(pingErr)
	}
	log.Println("Connection successful to the mysql db")
	sqlRows, sqlErr := Database.Query("select * from ar_internal_metadata;") // example for querying data and checking connection
	if sqlErr != nil{
		log.Fatal(sqlErr)
	}
	log.Println(sqlRows.Columns())
}

func fetchEnvBasedDatabase(env string) (*DatabaseServer, error) {
	switch(env){
	case "development":
		return &DbConfig.Development, nil
	default:
		return &DatabaseServer{}, errors.New("database configuration not found for the requesting environment")
	}
}