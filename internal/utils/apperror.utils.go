package utils

import "errors"

var (
	ErrEmailAlreadyExist = errors.New("email already exist")
	ErrFieldEmpty        = errors.New("field can't be empty")
	ErrWrongEmailPass    = errors.New("wrong email or password")
	ErrLengthSlug        = errors.New("slug must be 3-50 length of characters")
	ErrReservedWord      = errors.New("slug can't use reserved word")
	ErrFormatSlug        = errors.New("slug only contains alphanumeric and hypens")
	ErrSlugTaken         = errors.New("slug already taken")
	ErrNoRows            = errors.New("no rows affected")
)
