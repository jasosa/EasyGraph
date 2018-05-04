package graphql

import (
	"errors"
)

// ErrFieldsAddedWithoutObject is raised when fields are added but no object is added to the query
var (
	ErrFieldsAddedWithoutObject = errors.New("An object should be added to a query before add any fields")
)
