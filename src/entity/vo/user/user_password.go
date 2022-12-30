package user

import (
	"unicode/utf8"

	"github.com/ha-zu/learn-clean-dev/src/entity/vo"
)

type Password string

const (
	PASSWORD_MIN_LEN = 12
)

func PasswordValidate(pw string) (*Password, error) {
	if pw == "" {
		return nil, vo.ErrEmptyValue
	}

	if utf8.RuneCountInString(pw) < PASSWORD_MIN_LEN {
		return nil, vo.ErrValueIsTooShort
	}

	if err := checkPasswordRune(pw); err != nil {
		return nil, err
	}

	uPW := Password(pw)
	return &uPW, nil
}

func ChangePassword(pw Password) (*Password, error) {

	ckPW := string(pw)

	if ckPW == "" {
		return nil, vo.ErrEmptyValue
	}

	if utf8.RuneCountInString(ckPW) < PASSWORD_MIN_LEN {
		return nil, vo.ErrValueIsTooShort
	}

	if err := checkPasswordRune(ckPW); err != nil {
		return nil, err
	}

	uPW := Password(ckPW)

	return &uPW, nil
}

func checkPasswordRune(pw string) error {
	// ToDo check rune
	// for _, r := range pw {
	// }
	return nil
}