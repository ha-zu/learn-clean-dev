package user

import custerr "github.com/ha-zu/learn-clean-dev/src/entity/customerror"

type UserID string

func NewUserID(id UserID) error {

	if id == "" {
		return custerr.ErrEmptyValue
	}

	return nil

}
