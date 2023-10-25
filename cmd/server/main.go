package main

import (
	"errors"
	"github.com/EvgeniyBudaev/golang-family-mart/internal/auth"
	"github.com/EvgeniyBudaev/golang-family-mart/internal/register"
	"github.com/EvgeniyBudaev/golang-family-mart/internal/root"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	log.Println("Server starts...")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	LISTEN := os.Getenv("LISTEN")
	if LISTEN == "" {

	}
	log.Println(LISTEN)
	mux := http.NewServeMux()

	mux.HandleFunc("/", root.GetRoot)
	mux.HandleFunc("/api/v1/auth", auth.Auth)
	mux.HandleFunc("/api/v1/register", register.Register)

	err = http.ListenAndServe(LISTEN, mux)
	if errors.Is(err, http.ErrServerClosed) {
		log.Println("server closed")
	} else if err != nil {
		log.Panic("error starting server: %s", err)
	}
}
