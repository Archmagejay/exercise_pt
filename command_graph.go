package main

import (
	"errors"
)

func commandGraph(s *state, args ...string) error {
	if len(args) == 0 {
		return ErrArgs
	}
	return errors.New("not Implemented")
}