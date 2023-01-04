package mentor

import custerr "github.com/ha-zu/learn-clean-dev/src/entity/customerror"

type MentorPlanID string

func NewMentorPlanID(id MentorPlanID) error {

	if id == "" {
		return custerr.ErrEmptyValue
	}

	return nil

}
