package tag

import custerr "github.com/ha-zu/learn-clean-dev/src/entity/customerror"

type TagID string

func NewTagID(id TagID) error {

	if id == "" {
		return custerr.ErrEmptyValue
	}

	return nil

}
