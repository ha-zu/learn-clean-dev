package category

import custerr "github.com/ha-zu/learn-clean-dev/src/entity/customerror"

type CategoryName string

func NewCategoryName(name string) (CategoryName, error) {

	if name == "" {
		return CategoryName(""), custerr.ErrEmptyValue
	}

	return CategoryName(name), nil

}
