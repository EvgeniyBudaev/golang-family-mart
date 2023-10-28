package me

import (
	"github.com/EvgeniyBudaev/golang-family-mart/internal/scope"
	"log"
	"net/http"
)

func Me(scope scope.Scope) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Me")
	}
}
