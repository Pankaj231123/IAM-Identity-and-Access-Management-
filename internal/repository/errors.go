package repository

import "errors"

var (
	ErrNotFound   = errors.New("record not found")
	ErrEmailTaken = errors.New("email already taken")
)
