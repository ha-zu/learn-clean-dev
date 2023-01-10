package user

import (
	"regexp"
	"unicode/utf8"

	custerr "github.com/ha-zu/learn-clean-dev/src/entity/customerror"
)

type Password string

const (
	PASSWORD_MIN_LEN = 12
	PASSWORD_REGEXP  = `^[a-zA-Z0-9]{12,}$`
)

func NewPassword(pw string) (Password, error) {
	return ValidatePassword(pw)
}

func ValidatePassword(pw string) (Password, error) {

	if pw == "" {
		return Password(""), custerr.ErrEmptyValue
	}

	if utf8.RuneCountInString(pw) < PASSWORD_MIN_LEN {
		return Password(pw), custerr.ErrValueIsTooShort
	}

	r := regexp.MustCompile(PASSWORD_REGEXP)
	if !r.MatchString(pw) {
		return Password(pw), custerr.ErrNotCorrectFormat
	}

	return Password(pw), nil

}
