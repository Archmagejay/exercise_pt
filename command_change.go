package main

import (
	"errors"
)

func commandChange(s *state, args ...string) error {
	if len(args) == 0 {
		return ErrArgs
	}
	return errors.New("not Implemented")
}
