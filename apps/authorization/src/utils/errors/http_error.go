package errors

import (
	"encoding/json"
)

func GetErrorBody(err error) (ErrorBody, []byte) {

	errorBody := ErrorBody{}
	newError, _ := json.Marshal(err)
	json.Unmarshal([]byte(newError), &errorBody)

	return errorBody, newError
}
