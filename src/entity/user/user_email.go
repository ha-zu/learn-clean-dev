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

func NewEmail(email Email) error {

	if email == "" {
		return custerr.ErrEmptyValue
	}

	chkEmail := string(email)

	if l := utf8.RuneCountInString(chkEmail); l > EMAIL_MAX_LEN {
		return custerr.ErrValueIsTooLong
	}

	r := regexp.MustCompile(EMAIL_REGEXP)
	if !r.MatchString(chkEmail) {
		return custerr.ErrNotCorrectFormat
	}

	return nil
}

func ChangeEmail(email Email) error {

	if email == "" {
		return custerr.ErrEmptyValue
	}

	chkEmail := string(email)

	if l := utf8.RuneCountInString(chkEmail); l > EMAIL_MAX_LEN {
		return custerr.ErrValueIsTooLong
	}

	r := regexp.MustCompile(EMAIL_REGEXP)
	if !r.MatchString(chkEmail) {
		return custerr.ErrNotCorrectFormat
	}

	return nil

}
