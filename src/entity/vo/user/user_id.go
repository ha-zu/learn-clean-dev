package user

import "github.com/ha-zu/learn-clean-dev/src/entity/vo"

type UserID string

func UserIDValidate(id string) (*UserID, error) {
	// UserID Logic
	if id == "" {
		return nil, vo.ErrEmptyValue
	}

	uID := UserID(id)
	return &uID, nil
}
