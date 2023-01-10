package user

import (
	"github.com/google/uuid"
	custerr "github.com/ha-zu/learn-clean-dev/src/entity/customerror"
)

type UserID string

func NewUserID() UserID {
	return UserID(uuid.New().String())
}

func NewUserIDStr(id UserID) (UserID, error) {

	if id == "" {
		return UserID(""), custerr.ErrEmptyValue
	}

	return UserID(id), nil

}
