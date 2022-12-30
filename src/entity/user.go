package entity

import (
	"unicode/utf8"

	"github.com/ha-zu/learn-clean-dev/src/entity/vo"
	"github.com/ha-zu/learn-clean-dev/src/entity/vo/user"
)

type User struct {
	id       user.UserID
	userName string
	email    user.Email
	password user.Password
	profile  string
	skill    []Skill
	resume   []Resume
}

const (
	USER_NAME_MAX_LEN    = 255
	USER_PROFILE_MAX_LEN = 2000
)

func NewUser(id, name, email, password, profile string, skill []Skill, resume []Resume) (*User, error) {

	uID, err := user.UserIDValidate(id)
	if err != nil {
		return nil, err
	}

	if name == "" {
		return nil, vo.ErrEmptyValue
	}

	if l := utf8.RuneCountInString(name); l > USER_NAME_MAX_LEN {
		return nil, vo.ErrValueIsTooLong
	}

	uEmail, err := user.EmailValidate(email)
	if err != nil {
		return nil, err
	}

	uPassword, err := user.PasswordValidate(password)
	if err != nil {
		return nil, err
	}

	if l := utf8.RuneCountInString(profile); l > USER_PROFILE_MAX_LEN {
		return nil, vo.ErrValueIsTooLong
	}

	return &User{
		id:       *uID,
		userName: name,
		email:    *uEmail,
		password: *uPassword,
		profile:  profile,
		skill:    skill,
		resume:   resume,
	}, nil
}

func (u *User) GetUserID() user.UserID {
	return u.id
}

func (u *User) ChangeName(userName string) error {

	if userName == "" {
		return vo.ErrEmptyValue
	}

	if l := utf8.RuneCountInString(userName); l > USER_NAME_MAX_LEN {
		return vo.ErrValueIsTooLong
	}

	u.userName = userName

	return nil
}

func (u *User) ChangeProfile(profile string) error {

	if profile == "" {
		return vo.ErrEmptyValue
	}

	if l := utf8.RuneCountInString(profile); l > USER_PROFILE_MAX_LEN {
		return vo.ErrValueIsTooLong
	}

	u.profile = profile

	return nil
}
