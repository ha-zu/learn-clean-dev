package user

import custerr "github.com/ha-zu/learn-clean-dev/src/entity/customerror"

type ResumeID string

func NewResumeID(id ResumeID) error {

	if id == "" {
		return custerr.ErrEmptyValue
	}

	return nil
}
