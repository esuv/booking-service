package domain

import "errors"

var (
	ErrRequired = errors.New("required value")
	ErrNil      = errors.New("nil data")
)
