package user

import (
	"unicode/utf8"

	custerr "github.com/ha-zu/learn-clean-dev/src/entity/customerror"
)

type UserName string

const USER_NAME_MAX_LEN = 255

func NewUserName(name UserName) error {

	chkName := string(name)

	if chkName == "" {
		return custerr.ErrEmptyValue
	}

	if l := utf8.RuneCountInString(chkName); l > USER_NAME_MAX_LEN {
		return custerr.ErrValueIsTooLong
	}

	return nil
}

func ChangeName(name UserName) error {

	if name == "" {
		return custerr.ErrEmptyValue
	}

	chkName := string(name)

	if l := utf8.RuneCountInString(chkName); l > USER_NAME_MAX_LEN {
		return custerr.ErrValueIsTooLong
	}

	return nil

}
