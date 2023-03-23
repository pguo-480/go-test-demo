package pkg

import (
	"errors"
)

var ErrorRecoverable error = errors.New("recoverable error")
var ErrorUnrecoverable error = errors.New("unrecoverable error")
var ErrorDb409 error = errors.New("db 409")
