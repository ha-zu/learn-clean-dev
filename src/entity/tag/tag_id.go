package tag

import custerr "github.com/ha-zu/learn-clean-dev/src/entity/customerror"

type TagID string

func TagIDValid(id string) (*TagID, error) {

	if id == "" {
		return nil, custerr.ErrEmptyValue
	}

	tID := TagID(id)

	return &tID, nil
}
