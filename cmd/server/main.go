package main

import (
	"errors"
	"github.com/EvgeniyBudaev/golang-family-mart/internal/auth"
	"github.com/EvgeniyBudaev/golang-family-mart/internal/register"
	"github.com/EvgeniyBudaev/golang-family-mart/internal/root"
	"github.com/EvgeniyBudaev/golang-family-mart/internal/scope"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	scope := scope.NewScope()
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	log.Println("Connect database...")

	db, err := sqlx.Connect(os.Getenv("DATABASE_DRIVER"), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalln(err)
	}
	scope.Database = db
	log.Println("Server starts...")

	LISTEN := os.Getenv("LISTEN")
	if LISTEN == "" {

	}

	log.Println(LISTEN)
	mux := http.NewServeMux()

	mux.HandleFunc("/", root.GetRoot)
	mux.HandleFunc("/api/v1/auth", auth.Auth(scope))
	mux.HandleFunc("/api/v1/register", register.Register(scope))

	err = http.ListenAndServe(LISTEN, mux)
	if errors.Is(err, http.ErrServerClosed) {
		log.Println("server closed")
	} else if err != nil {
		log.Panic("error starting server: %s", err)
	}
}
