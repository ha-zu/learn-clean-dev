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

func NewPassword(pw Password) error {
	if pw == "" {
		return custerr.ErrEmptyValue
	}

	chkPW := string(pw)

	if utf8.RuneCountInString(chkPW) < PASSWORD_MIN_LEN {
		return custerr.ErrValueIsTooShort
	}

	if err := checkPassword(chkPW); err != nil {
		return err
	}

	return nil

}

func ChangePassword(pw Password) error {

	chkPW := string(pw)

	if chkPW == "" {
		return custerr.ErrEmptyValue
	}

	if utf8.RuneCountInString(chkPW) < PASSWORD_MIN_LEN {
		return custerr.ErrValueIsTooShort
	}

	if err := checkPassword(chkPW); err != nil {
		return err
	}

	return nil
}

func checkPassword(pw string) error {

	r := regexp.MustCompile(PASSWORD_REGEXP)

	if !r.MatchString(pw) {
		return custerr.ErrNotCorrectFormat
	}

	return nil
}
