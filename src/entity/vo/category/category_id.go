package category

import "github.com/ha-zu/learn-clean-dev/src/entity/vo"

type CategoryID string

func CategoryIDValidate(id string) (*CategoryID, error) {

	if id == "" {
		return nil, vo.ErrEmptyValue
	}

	cID := CategoryID(id)

	return &cID, nil
}
