package errors_test

import (
	"fmt"
	"testing"

	"github.com/magiconair/properties/assert"
	"github.com/mviniciusgc/authorization/src/utils/errors"
)

func TestError_Error(t *testing.T) {
	e := errors.Error{}
	assert.Equal(t, "", e.Error())
	e.Code = errors.EFORBIDDEN
	assert.Equal(t, "<forbidden> ", e.Error())
	e.Message = "error message!"
	assert.Equal(t, "<forbidden> error message!", e.Error())
	e.Op = "user.services.Create"
	assert.Equal(t, "user.services.Create: <forbidden> error message!", e.Error())
	e.Detail = []byte("abacate")
	assert.Equal(t, "user.services.Create: <forbidden> error message!", e.Error())
	e.Err = &errors.Error{
		Code:    errors.EFORBIDDEN,
		Message: "embed error message!",
		Op:      "user.repository.Create",
		Detail:  []byte(`{"value":"abacate"}`),
	}
	assert.Equal(t, "user.services.Create: user.repository.Create: <forbidden> embed error message!", e.Error())
}

func TestError_DetailedError(t *testing.T) {
	e := errors.Error{}
	assert.Equal(t, "", e.DetailedError())
	e.Code = errors.EFORBIDDEN
	assert.Equal(t, "Code: forbidden\n", e.DetailedError())
	e.Message = "error message!"
	assert.Equal(t, "Code: forbidden\nMessage: error message!\n", e.DetailedError())
	e.Op = "user.services.Create"
	assert.Equal(t, "Op: user.services.Create\nCode: forbidden\nMessage: error message!\n", e.DetailedError())
	e.Detail = []byte("abacate")
	assert.Equal(t, "Op: user.services.Create\nCode: forbidden\nMessage: error message!\nDetail: abacate\n", e.DetailedError())
	e.Err = &errors.Error{
		Code:    errors.EFORBIDDEN,
		Message: "embed error message!",
		Op:      "user.repository.Create",
		Detail:  []byte(`{"value":"abacate"}`),
	}
	assert.Equal(t, "Op: user.services.Create\nCode: forbidden\nMessage: error message!\nDetail: abacate\nErr:\n\tOp: user.repository.Create\n\tCode: forbidden\n\tMessage: embed error message!\n\tDetail:\n\t{\n\t\t\"value\": \"abacate\"\n\t}", e.DetailedError())
	e.Err = &errors.Error{
		Code:    errors.EFORBIDDEN,
		Message: "embed error message!",
		Op:      "user.repository.Create",
		Detail:  []byte(`{"value":"abacate"}`),
		Err:     fmt.Errorf("embed error message!"),
	}
	assert.Equal(t, "Op: user.services.Create\nCode: forbidden\nMessage: error message!\nDetail: abacate\nErr:\n\tOp: user.repository.Create\n\tCode: forbidden\n\tMessage: embed error message!\n\tDetail:\n\t{\n\t\t\"value\": \"abacate\"\n\t}\tErr: embed error message!\n", e.DetailedError())
}

func TestError_ErrorCode(t *testing.T) {
	assert.Equal(t, "", errors.ErrorCode(nil))

	e := errors.Error{}
	assert.Equal(t, "internal", errors.ErrorCode(&e))

	e = errors.Error{Err: &errors.Error{Code: errors.EBADREQUEST}}
	assert.Equal(t, "bad_request", errors.ErrorCode(&e))

	e = errors.Error{Code: errors.EFORBIDDEN}
	assert.Equal(t, "forbidden", errors.ErrorCode(&e))
}

func TestError_ErrorMessage(t *testing.T) {
	assert.Equal(t, "", errors.ErrorMessage(nil))

	e := errors.Error{}
	assert.Equal(t, "An internal error has occurred. Please contact technical support.", errors.ErrorMessage(&e))

	e = errors.Error{Err: &errors.Error{Message: "Custom message"}}
	assert.Equal(t, "Custom message", errors.ErrorMessage(&e))

	e = errors.Error{Message: "Custom message 2"}
	assert.Equal(t, "Custom message 2", errors.ErrorMessage(&e))
}

func TestError_Stacktrace(t *testing.T) {
	e := errors.Error{Err: &errors.Error{Code: errors.EBADREQUEST}}
	stack := fmt.Sprint(e.StackTrace())
	assert.Equal(t, "[]", stack)
}
