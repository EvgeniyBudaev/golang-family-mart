package internal_errors

import (
	"encoding/json"
	"github.com/EvgeniyBudaev/golang-family-mart/internal/schemas"
	"log"
	"net/http"
)

func ErrorJson(w http.ResponseWriter, err error, status int) {
	log.Panicln("Error:", err)
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	jsonSerialized, jsonErr := json.Marshal(schemas.NewErrorResponse(err))
	if jsonErr != nil {
		jsonSerialized = []byte("{\"error\":\"Cant serilize error to json\"}")
	}
	w.Write(jsonSerialized)
}
