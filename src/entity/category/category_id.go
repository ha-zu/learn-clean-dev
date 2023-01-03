package category

import custerr "github.com/ha-zu/learn-clean-dev/src/entity/customerror"

type CategoryID string

func NewCategoryID(id CategoryID) error {

	if id == "" {
		return custerr.ErrEmptyValue
	}

	return nil
}
