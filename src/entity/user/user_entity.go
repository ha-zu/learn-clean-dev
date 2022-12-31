package user

import (
	"unicode/utf8"

	custerr "github.com/ha-zu/learn-clean-dev/src/entity/customerror"
)

type User struct {
	id       UserID
	userName string
	email    Email
	password Password
	profile  string
	skill    []Skill
	resume   []Resume
}

const (
	USER_NAME_MAX_LEN    = 255
	USER_PROFILE_MAX_LEN = 2000
)

func NewUser(id, name, email, password, profile string, skill []Skill, resume []Resume) (*User, error) {

	uID, err := UserIDValidate(id)
	if err != nil {
		return nil, err
	}

	if name == "" {
		return nil, custerr.ErrEmptyValue
	}

	if l := utf8.RuneCountInString(name); l > USER_NAME_MAX_LEN {
		return nil, custerr.ErrValueIsTooLong
	}

	uEmail, err := EmailValidate(email)
	if err != nil {
		return nil, err
	}

	uPassword, err := PasswordValidate(password)
	if err != nil {
		return nil, err
	}

	if l := utf8.RuneCountInString(profile); l > USER_PROFILE_MAX_LEN {
		return nil, custerr.ErrValueIsTooLong
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

func (u *User) GetUserID() UserID {
	return u.id
}

func (u *User) GetIsProposal() bool {
	return len(u.skill) < 5
}

func (u *User) ChangeName(userName string) error {

	if userName == "" {
		return custerr.ErrEmptyValue
	}

	if l := utf8.RuneCountInString(userName); l > USER_NAME_MAX_LEN {
		return custerr.ErrValueIsTooLong
	}

	u.userName = userName

	return nil
}

func (u *User) ChangeProfile(profile string) error {

	if profile == "" {
		return custerr.ErrEmptyValue
	}

	if l := utf8.RuneCountInString(profile); l > USER_PROFILE_MAX_LEN {
		return custerr.ErrValueIsTooLong
	}

	u.profile = profile

	return nil
}
