package utils

import (
	"testing"

	"github.com/Nerzal/gocloak/v13"
	"github.com/mviniciusgc/authorization/src/utils/errors"
	"github.com/stretchr/testify/assert"
)

func TestVerifyErrors(t *testing.T) {
	t.Run("Return error 400", func(t *testing.T) {
		err := &gocloak.APIError{Code: 400}
		errorResponse := VerifyErrors(err, "Error EBADREQUEST")
		assert.NotNil(t, errorResponse)
		assert.Equal(t, "bad_request", errorResponse.(*errors.Error).Code)
	})
	t.Run("Return error 409", func(t *testing.T) {
		err := &gocloak.APIError{Code: 409}
		errorResponse := VerifyErrors(err, "Error ECONFLICT")
		assert.NotNil(t, errorResponse)
		assert.Equal(t, "conflict", errorResponse.(*errors.Error).Code)
	})
	t.Run("Return error 500", func(t *testing.T) {
		err := &gocloak.APIError{Code: 500}
		errorResponse := VerifyErrors(err, "Error EINTERNAL")
		assert.NotNil(t, errorResponse)
		assert.Equal(t, "internal", errorResponse.(*errors.Error).Code)
	})

	t.Run("Return error default", func(t *testing.T) {
		err := &gocloak.APIError{Code: 300}
		errorResponse := VerifyErrors(err, "Error EINTERNAL")
		assert.NotNil(t, errorResponse)
		assert.Equal(t, "internal", errorResponse.(*errors.Error).Code)
	})

}
