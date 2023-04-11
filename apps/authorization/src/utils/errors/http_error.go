package errors

import (
	"encoding/json"
	"io"
)

type HttpError struct {
	Message string `json:"message,omitempty"`
}

func FormatHttpError(body io.Reader) *Error {
	httpError := HttpError{}
	err := json.NewDecoder(body).Decode(&httpError)
	if err != nil {
		return &Error{Op: "FormatHttpError.onAuth0", Err: err}
	}

	return &Error{Message: httpError.Message}
}

func GetErrorBody(err error) (ErrorBody, []byte) {
	errorBody := ErrorBody{}
	newError, _ := json.Marshal(err)
	json.Unmarshal([]byte(newError), &errorBody)

	return errorBody, newError
}
