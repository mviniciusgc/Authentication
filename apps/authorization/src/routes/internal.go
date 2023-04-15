package routes

import (
	"encoding/json"
	"net/http"

	"github.com/mviniciusgc/authorization/src/entity"
	"github.com/mviniciusgc/authorization/src/utils/errors"
)

func createUser(s *HandlerServices) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		user := &entity.UserRequest{}
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		resp, err := s.KeycloakController.CreateUser(*user)
		if err != nil {
			err, errByte := errors.GetErrorBody(err)
			w.WriteHeader(*err.Err.Code)
			w.Write(errByte)
			return
		}

		userResp, _ := json.Marshal(resp)

		w.WriteHeader(http.StatusCreated)
		w.Write(userResp)
	})
}
