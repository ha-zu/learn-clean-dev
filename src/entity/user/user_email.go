package user

import (
	"regexp"
	"unicode/utf8"

	custerr "github.com/ha-zu/learn-clean-dev/src/entity/customerror"
)

type Email string

const (
	EMAIL_MAX_LEN = 255
	EMAIL_REGEXP  = `^([a-z0-9]+[a-z0-9.+-_][a-z0-9])+@([a-z0-9]+[a-z0-9.+-_]*\.)[a-z]{2,}$`
)

func EmailValidate(email string) (*Email, error) {

	if email == "" {
		return nil, custerr.ErrEmptyValue
	}

	if l := utf8.RuneCountInString(email); l > EMAIL_MAX_LEN {
		return nil, custerr.ErrValueIsTooLong
	}

	r := regexp.MustCompile(EMAIL_REGEXP)
	if !r.MatchString(email) {
		return nil, custerr.ErrNotCorrectFormat
	}

	uEmail := Email(email)
	return &uEmail, nil
}
