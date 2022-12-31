package mentor

import custerr "github.com/ha-zu/learn-clean-dev/src/entity/customerror"

type MentorPlanID string

func MentorPlanIDValidate(id string) (*MentorPlanID, error) {

	if id == "" {
		return nil, custerr.ErrEmptyValue
	}

	mID := MentorPlanID(id)

	return &mID, nil
}
