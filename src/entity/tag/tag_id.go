package tag

import (
	"github.com/google/uuid"
	custerr "github.com/ha-zu/learn-clean-dev/src/entity/customerror"
)

type TagID string

func NewTagID() TagID {
	return TagID(uuid.New().String())
}

func NewTagIDStr(id TagID) (TagID, error) {

	if id == "" {
		return TagID(""), custerr.ErrEmptyValue
	}

	return TagID(id), nil

}
