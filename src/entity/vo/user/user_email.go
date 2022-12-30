package user

import (
	"regexp"
	"unicode/utf8"

	"github.com/ha-zu/learn-clean-dev/src/entity/vo"
)

type Email string

const (
	EMAIL_MAX_LEN = 255
	//ToDo Regexp
	EMAIL_REGEXP = ``
)

func EmailValidate(email string) (*Email, error) {

	if email == "" {
		return nil, vo.ErrEmptyValue
	}

	if l := utf8.RuneCountInString(email); l > EMAIL_MAX_LEN {
		return nil, vo.ErrValueIsTooLong
	}

	r := regexp.MustCompile(EMAIL_REGEXP)
	if !r.MatchString(email) {
		return nil, vo.ErrNotCorrectFormat
	}

	uEmail := Email(email)
	return &uEmail, nil
}
