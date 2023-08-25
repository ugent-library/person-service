package models

import "errors"

var ErrNotFound = errors.New("not found")
var ErrFatal = errors.New("fatal error")
var ErrNonFatal = errors.New("non fatal error")
var ErrSkipped = errors.New("was skipped")
var ErrMissingArgument = errors.New("missing argument")

type Repository interface {
	PersonService
	PersonSuggestService
	OrganizationService
	OrganizationSuggestService
}
