package exceptions

import "errors"

func NewInvalidArgument(message string) error {
	return errors.New(message)
}
