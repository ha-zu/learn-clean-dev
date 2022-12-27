package vo

import "errors"

type User struct {
	id       string
	userName string
	email    string
	password string
	profile  string
}

func UserVerify(id, userName, email, password, profile string) (*User, error) {

	if id == "" {
		return nil, errors.New("must not be empty")
	}

	if userName == "" {
		return nil, errors.New("must not be empty")
	}

	if email == "" {
		return nil, errors.New("must not be empty")
	}

	if password == "" {
		return nil, errors.New("must not be empty")
	}

	//ToDo Validation

	return &User{
		id:       id,
		userName: userName,
		email:    email,
		password: password,
		profile:  profile,
	}, nil
}
