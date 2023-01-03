package category

import custerr "github.com/ha-zu/learn-clean-dev/src/entity/customerror"

type CategoryName string

func NewCategoryName(name CategoryName) error {

	if name == "" {
		return custerr.ErrEmptyValue
	}

	return nil

}
