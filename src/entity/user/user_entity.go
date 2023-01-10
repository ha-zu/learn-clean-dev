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

	u := &User{
		id:       id,
		name:     name,
		email:    email,
		password: password,
		skill:    skill,
		resume:   resume,
	}

	err := u.ValidateProfile(profile)
	if err != nil {
		return nil, err
	}

	return u, nil

}

func (u *User) GetUserID() UserID {
	return u.id
}

func (u *User) GetIsProposal() bool {
	return len(u.skill) < 5
}

func (u *User) ValidateProfile(profile string) error {

	if l := utf8.RuneCountInString(profile); l > USER_PROFILE_MAX_LEN {
		return custerr.ErrValueIsTooLong
	}

	u.profile = profile

	return nil

}
