package routes

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mviniciusgc/authorization/src/entity"
	"github.com/mviniciusgc/authorization/src/utils/errors"
)

func UpdateUser(s *HandlerServices) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		authHeader := r.Header.Get("Authorization")
		err := s.middleware.ParseToken(authHeader)
		if err != nil {
			_, errByte := errors.GetErrorBody(err)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write(errByte)
			return
		}

		userID := chi.URLParam(r, "userID")
		if userID == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		user := &entity.UserUpdateRequest{}
		err = json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		err = s.KeycloakController.UpdateUser(*user, userID)
		if err != nil {
			err, errByte := errors.GetErrorBody(err)
			w.WriteHeader(*err.Err.Code)
			w.Write(errByte)
			return
		}

		w.WriteHeader(http.StatusOK)
	})
}
