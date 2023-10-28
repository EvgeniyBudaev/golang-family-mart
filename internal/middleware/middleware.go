package middleware

import (
	"github.com/gorilla/handlers"
	"github.com/justinas/alice"
	"log"
	"net/http"
	"os"
)

func Jwt(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Print("Executing middlewareTwo")
		if r.URL.Path == "/foo" {
			return
		}

		next.ServeHTTP(w, r)
		log.Print("Executing middlewareTwo again")
	})
}
func NewLoggingHandler() func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		logFile, err := os.OpenFile("server.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0664)
		if err != nil {
			log.Fatal("NewLoggingHandler:", err)
		}
		return handlers.LoggingHandler(logFile, h)
	}
}

func AuthedApiChain() alice.Chain {
	return alice.New(NewLoggingHandler(), Jwt)
}
