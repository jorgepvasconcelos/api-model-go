package bookError

import (
	"errors"
)

type BookError error

var (
	NotFound BookError = errors.New("")
)
