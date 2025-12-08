package main

import (
	"errors"
)

func commandExport(s *state, args ...string) error {

	return errors.New("not Implemented")
}

func commandImport(s *state, args ...string) error {
	if len(args) == 0 {
		return ErrArgs
	}
	return errors.New("not Implemented")
}