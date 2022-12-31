package category

import custerr "github.com/ha-zu/learn-clean-dev/src/entity/customerror"

type CategoryID string

func CategoryIDValidate(id string) (*CategoryID, error) {

	if id == "" {
		return nil, custerr.ErrEmptyValue
	}

	cID := CategoryID(id)

	return &cID, nil
}
