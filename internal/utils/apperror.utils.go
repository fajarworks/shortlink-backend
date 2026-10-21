package utils

import "errors"

var (
	ErrEmailAlreadyExist = errors.New("email already exist")
	ErrFieldEmpty        = errors.New("field can't be empty")
	ErrWrongEmailPass    = errors.New("wrong email or password")
)
