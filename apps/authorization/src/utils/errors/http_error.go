package errors

import (
	"encoding/json"
)

type HttpError struct {
	Message string `json:"message,omitempty"`
}

func GetErrorBody(err error) (ErrorBody, []byte) {
	errorBody := ErrorBody{}
	newError, _ := json.Marshal(err)
	json.Unmarshal([]byte(newError), &errorBody)

	return errorBody, newError
}
