package sqlError

import (
	"errors"
)

type SqlError error

var (
	NotFound SqlError = errors.New("")
)
