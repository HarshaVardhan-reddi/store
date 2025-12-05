package helpers

import (
	"encoding/json"
	"net/http"
)

type ResponseConstructor struct{
	Status int
	Message string
	Object any
	Writer http.ResponseWriter
}

func (rc *ResponseConstructor) RenderJSON() {
	response := make(map[string]any)
	rc.Writer.Header().Add("Content-Type","application/json")
	rc.Writer.WriteHeader(rc.Status)

	response["message"] = rc.Message
	response["data"] = rc.Object

	if err := json.NewEncoder(rc.Writer).Encode(response); err != nil {
		http.Error(rc.Writer, `{"message":"failed to encode response"}`, http.StatusInternalServerError)
	}
}