package register

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/EvgeniyBudaev/golang-family-mart/internal/internal_errors"
	"github.com/EvgeniyBudaev/golang-family-mart/internal/models"
	"github.com/EvgeniyBudaev/golang-family-mart/internal/schemas"
	"github.com/EvgeniyBudaev/golang-family-mart/internal/scope"
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
		if err == sql.ErrNoRows {
			internal_errors.ErrorJson(w, errors.New("user allready exists"), http.StatusBadRequest)
			return
		}

	}
}
