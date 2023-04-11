package middleware

import (
	"github.com/Nerzal/gocloak/v13"
	"github.com/mviniciusgc/authorization/src/utils/errors"
)

func VerifyErrors(err error, op string) error {
	switch err := err.(type) {
	case *gocloak.APIError:
		switch err.Code {
		case 400:
			return errors.NewError(&errors.Error{Op: op, Err: err, Code: errors.EBADREQUEST})
		case 409:
			return errors.NewError(&errors.Error{Op: op, Err: err, Code: errors.ECONFLICT})
		case 500:
			return errors.NewError(&errors.Error{Op: op, Err: err, Code: errors.EINTERNAL})
		default:
			return errors.NewError(&errors.Error{Op: op, Err: err, Code: errors.EINTERNAL})
		}
	}

	return nil
}
