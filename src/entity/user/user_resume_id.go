package user

import custerr "github.com/ha-zu/learn-clean-dev/src/entity/customerror"

type ResumeID string

func ResumeIDValid(id string) (*ResumeID, error) {

	if id == "" {
		return nil, custerr.ErrEmptyValue
	}

	rID := ResumeID(id)

	return &rID, nil
}
