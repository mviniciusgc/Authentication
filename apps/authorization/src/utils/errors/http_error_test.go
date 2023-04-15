package errors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetErrorBody(t *testing.T) {
	t.Run("Not return error", func(t *testing.T) {
		err := NewError(&Error{Code: EBADREQUEST})
		errorBody, newError := GetErrorBody(err)
		assert.NotNil(t, errorBody)
		assert.NotNil(t, newError)
	})

}
