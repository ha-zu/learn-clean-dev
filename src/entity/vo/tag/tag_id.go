package tag

import "github.com/ha-zu/learn-clean-dev/src/entity/vo"

type TagID string

func TagIDValid(id string) (*TagID, error) {

	if id == "" {
		return nil, vo.ErrEmptyValue
	}

	tID := TagID(id)

	return &tID, nil
}
