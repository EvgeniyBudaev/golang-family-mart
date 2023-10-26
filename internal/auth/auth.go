package auth

import (
	"github.com/EvgeniyBudaev/golang-family-mart/internal/scope"
	"log"
	"net/http"
)

func Auth(scope scope.Scope) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(scope)
		log.Println(r)
	}
}
