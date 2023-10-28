package auth

import (
	"encoding/json"
	"github.com/EvgeniyBudaev/golang-family-mart/internal/scope"
	"log"
	"net/http"
)

type test_struct struct {
	Test string
}

func Auth(scope scope.Scope) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var t test_struct
		err := decoder.Decode(&t)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		log.Println(t.Test)
	}
}
