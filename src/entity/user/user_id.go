package user

import custerr "github.com/ha-zu/learn-clean-dev/src/entity/customerror"

type UserID string

func UserIDValidate(id string) (*UserID, error) {

	if id == "" {
		return nil, custerr.ErrEmptyValue
	}

	uID := UserID(id)
	return &uID, nil
}
