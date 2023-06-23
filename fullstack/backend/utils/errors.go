package utils

import "errors"

var (
	ErrNotFound         = errors.New("not found")
	ErrUnauthorized     = errors.New("unauthorized")
	ErrInternalServerError = errors.New("internal server error")
)
