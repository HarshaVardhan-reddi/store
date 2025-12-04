package helpers

import (
	"encoding/json"
	"net/http"
)

type RepsonseConstructor struct{
	Status int
	Message string
	Object any
	Writer http.ResponseWriter
}

func (rc *RepsonseConstructor) RenderJSON() {
	var response any
	rc.Writer.Header().Add("Content-Type","application/json")
	rc.Writer.WriteHeader(rc.Status)

	if rc.Message != ""{
		response = map[string]string{"message":rc.Message}
	}else if rc.Object != nil{
		response = rc.Object
	}

	if err := json.NewEncoder(rc.Writer).Encode(response); err != nil {
		http.Error(rc.Writer, `{"message":"failed to encode response"}`, http.StatusInternalServerError)
	}
}