package user

import (
	"github.com/google/uuid"
	custerr "github.com/ha-zu/learn-clean-dev/src/entity/customerror"
)

type SkillID string

func NewSkillID() SkillID {
	return SkillID(uuid.New().String())
}

func NewSkillIDStr(id string) (SkillID, error) {

	if id == "" {
		return SkillID(""), custerr.ErrEmptyValue
	}

	return SkillID(id), nil

}
