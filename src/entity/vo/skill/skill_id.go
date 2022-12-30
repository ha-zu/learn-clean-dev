package skill

import "github.com/ha-zu/learn-clean-dev/src/entity/vo"

type SkillID string

func SkillIDValid(id string) (*SkillID, error) {

	if id == "" {
		return nil, vo.ErrEmptyValue
	}

	sID := SkillID(id)

	return &sID, nil
}
