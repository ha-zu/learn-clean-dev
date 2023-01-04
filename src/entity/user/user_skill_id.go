package user

import custerr "github.com/ha-zu/learn-clean-dev/src/entity/customerror"

type SkillID string

func NewSkillID(id SkillID) error {

	if id == "" {
		return custerr.ErrEmptyValue
	}

	return nil

}
