package user

import custerr "github.com/ha-zu/learn-clean-dev/src/entity/customerror"

type SkillID string

func SkillIDValid(id string) (*SkillID, error) {

	if id == "" {
		return nil, custerr.ErrEmptyValue
	}

	sID := SkillID(id)

	return &sID, nil
}
