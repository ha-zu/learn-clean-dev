package user

import (
	"unicode/utf8"

	custerr "github.com/ha-zu/learn-clean-dev/src/entity/customerror"
)

type UserName string

const USER_NAME_MAX_LEN = 255

func NewUserName(name string) (UserName, error) {
	return ValidateUserName(name)
}

func ValidateUserName(name string) (UserName, error) {

	if name == "" {
		return UserName(""), custerr.ErrEmptyValue
	}

	if l := utf8.RuneCountInString(name); l > USER_NAME_MAX_LEN {
		return UserName(name), custerr.ErrValueIsTooLong
	}

	return UserName(name), nil

}
