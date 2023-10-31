package register

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/EvgeniyBudaev/golang-family-mart/internal/internal_errors"
	"github.com/EvgeniyBudaev/golang-family-mart/internal/internal_jwt"
	"github.com/EvgeniyBudaev/golang-family-mart/internal/models"
	"github.com/EvgeniyBudaev/golang-family-mart/internal/schemas"
	"github.com/EvgeniyBudaev/golang-family-mart/internal/scope"
	"log"
	"net/http"
)

func Register(scope scope.Scope) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var request schemas.RegisterRequest
		if err := decoder.Decode(&request); err != nil {
			internal_errors.ErrorJson(w, err, http.StatusBadRequest)
			return
		}
		_, err := models.GetUserByUsername(scope.Database, request.Username)
		log.Println(err)
		if err != sql.ErrNoRows {
			internal_errors.ErrorJson(w, errors.New("user allready exists"), http.StatusBadRequest)
			return
		}
		_, err = models.GetUserByEmail(scope.Database, request.Email)
		if err != sql.ErrNoRows {
			internal_errors.ErrorJson(w, errors.New("email allready used"), http.StatusBadRequest)
			return
		}
		if request.Password1 != request.Password2 {
			internal_errors.ErrorJson(w, errors.New("Password not same"), http.StatusBadRequest)
			return
		}
		user := models.User{
			Username: request.Username,
			Password: request.Password1,
			Email:    request.Email,
		}
		if user, err = models.AddUser(scope.Database, user); err != nil {
			internal_errors.ErrorJson(w, errors.New("Password not same"), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		token, err := internal_jwt.GenerateJWT(user)
		if err != nil {
			internal_errors.ErrorJson(w, err, http.StatusInternalServerError)
			return
		}
		jsonSerialized, _ := json.Marshal(schemas.AuthResponse{
			UserId:   user.Id,
			JwtToken: token,
		})
		w.Write(jsonSerialized)

		return
	}
}
