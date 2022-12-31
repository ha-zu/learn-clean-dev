package customerror

import "errors"

var (
	ErrEmptyValue       = errors.New("must be not empty value")
	ErrValueIsTooLong   = errors.New("parameters are too long")
	ErrValueIsTooShort  = errors.New("parameters are too short")
	ErrNotCorrectFormat = errors.New("not in the correct format")
	ErrOutOfRange       = errors.New("parameters are out of range")
	ErrNotEnoughValue   = errors.New("parameters not enough value")
	ErrNotDefaultValue  = errors.New("parameters not default value")
	ErrNotMatchValue    = errors.New("parameters not match value")
	ErrPerminssion      = errors.New("you do not have permission")
)
