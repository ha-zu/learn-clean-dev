package mentor

import "github.com/ha-zu/learn-clean-dev/src/entity/vo"

type MentorPlanID string

func MentorPlanIDValidate(id string) (*MentorPlanID, error) {

	if id == "" {
		return nil, vo.ErrEmptyValue
	}

	mID := MentorPlanID(id)

	return &mID, nil
}
