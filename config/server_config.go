package config

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	_ "embed"
)

const SERVER_CONF_FILE string = "config/serverconfig.json"

const SERVER_CONFIG_VALID = true
const SERVER_CONFIG_INVALID = false

// ports range
const SERVER_PORT_MIN_RANGE = 1024
const SERVER_PORT_MAX_RANGE = 49151

//go:embed serverconfig.json
var rawContent []byte
var ServerConfig ServerConfiguration

type ServerConfiguration struct{
	Port int32 `json:"port"`
	ReadTimeOut int8 `json:"read_timeout_in_secs"`
	WriteTimeOut int8 `json:"write_timeout_in_secs"`
	IdleTimeOut int8 `json:"idle_timeout_in_secs"`
}

func init(){
	// rawContent, err := os.ReadFile(SERVER_CONF_FILE)
	// if err != nil{
	// 	log.Fatal("Server configuration",err)
	// }
	if !json.Valid(rawContent) {
		log.Fatal("Server configuration file has invalid content")
	}

	errInParsing := json.Unmarshal(rawContent, &ServerConfig); if errInParsing != nil {
		log.Fatal("Server configuration",errInParsing)
	}

	ok, validationError := validateServerConfiguration(&ServerConfig)
	if(!ok){
		log.Fatal("Invalid Server configuration", validationError)
	}
	// config, _ := json.Marshal(ServerConfig)
	// log.Println(string(config))
	log.Println("\nserver configuration",string(rawContent))
	log.Println("Server configuration parsed successfully...")
}

func (serverConf ServerConfiguration) ConstructHttpServer(r *mux.Router) *http.Server {
	server := &http.Server{
		ReadTimeout: time.Duration(serverConf.ReadTimeOut) * time.Second,
		WriteTimeout: time.Duration(serverConf.WriteTimeOut) * time.Second,
		Addr: ":"+ strconv.Itoa(int(serverConf.Port)),
		IdleTimeout: time.Duration(serverConf.IdleTimeOut) * time.Second,
		Handler: r,
	}
	return server
}


func validateServerConfiguration(config *ServerConfiguration)(bool, error){
	if config.Port < SERVER_PORT_MIN_RANGE || config.Port > SERVER_PORT_MAX_RANGE {
		return false, errors.New("invalid port")
	}
	return SERVER_CONFIG_VALID,nil
}