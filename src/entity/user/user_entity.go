package user

import (
	"unicode/utf8"

	custerr "github.com/ha-zu/learn-clean-dev/src/entity/customerror"
)

type User struct {
	id       UserID
	name     UserName
	email    Email
	password Password
	profile  string
	skill    []Skill
	resume   []Resume
}

const USER_PROFILE_MAX_LEN = 2000

func NewUser(id UserID, name UserName, email Email, password Password, profile string, skill []Skill, resume []Resume) (*User, error) {

	err := NewUserID(id)
	if err != nil {
		return nil, err
	}

	err = NewUserName(name)
	if err != nil {
		return nil, err
	}

	err = NewEmail(email)
	if err != nil {
		return nil, err
	}

	err = NewPassword(password)
	if err != nil {
		return nil, err
	}

	if l := utf8.RuneCountInString(profile); l > USER_PROFILE_MAX_LEN {
		return nil, custerr.ErrValueIsTooLong
	}

	return &User{
		id:       id,
		name:     name,
		email:    email,
		password: password,
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

func (u *User) ChangeProfile(profile string) error {

	if l := utf8.RuneCountInString(profile); l > USER_PROFILE_MAX_LEN {
		return custerr.ErrValueIsTooLong
	}

	u.profile = profile

	return nil

}
