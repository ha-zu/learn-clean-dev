package resume

import "github.com/ha-zu/learn-clean-dev/src/entity/vo"

type ResumeID string

func ResumeIDValid(id string) (*ResumeID, error) {

	if id == "" {
		return nil, vo.ErrEmptyValue
	}

	rID := ResumeID(id)

	return &rID, nil
}
