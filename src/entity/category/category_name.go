package category

import custerr "github.com/ha-zu/learn-clean-dev/src/entity/customerror"

type CategoryName string

func CategoryNameValidate(name string) (*CategoryName, error) {

	if name == "" {
		return nil, custerr.ErrEmptyValue
	}

	custName := CategoryName(name)

	return &custName, nil

}
