package berror

import (
	"errors"
)

func ConversionError(err error) *ThrowError {
	if err == nil {
		return nil
	}
	var e *ThrowError
	if errors.As(err, &e) {
		return e
	}
	return &ThrowError{
		Err: err,
	}
}
