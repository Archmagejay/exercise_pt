package main

func commandExport(s *state, args ...string) error {
	return ErrNotImplemented
}

func commandImport(s *state, args ...string) error {
	if len(args) == 0 {
		return ErrArgs
	}
	return ErrNotImplemented
}
